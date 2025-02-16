package controllers

import (
	"influencer-golang/config"
	"influencer-golang/models"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

// Load .env saat aplikasi pertama kali berjalan
func init() {
	godotenv.Load()
	rand.Seed(time.Now().UnixNano()) // Seed rand agar Order ID unik
}

// Generate Order ID unik
func generateOrderID() string {
	return "ORDER-" + strconv.Itoa(rand.Intn(1000000))
}

// CreatePayment menangani proses pembuatan pembayaran
func CreatePayment(c *gin.Context) {
	var payment models.Payment

	// Validasi request body
	if err := c.ShouldBindJSON(&payment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validasi BuyerID dan ServiceID
	if payment.BuyerID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "BuyerID is required"})
		return
	}
	if payment.ServiceID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ServiceID is required"})
		return
	}

	// Cek apakah service tersedia
	var service models.Service
	if err := config.DB.First(&service, payment.ServiceID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Service not found"})
		return
	}

	// Cek apakah brand (buyer) tersedia
	var brand models.Brand
	if err := config.DB.First(&brand, payment.BuyerID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Brand not found"})
		return
	}

	// Generate Order ID dan set status pembayaran
	payment.TransactionID = generateOrderID()
	payment.Status = "pending"
	payment.Amount = service.PricePerPost

	// Simpan pembayaran ke database
	if err := config.DB.Create(&payment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create payment"})
		return
	}

	// Midtrans Payment Gateway
	snapGateway := snap.Client{}
	snapGateway.New(os.Getenv("MIDTRANS_SERVER_KEY"), midtrans.Sandbox)

	snapReq := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  payment.TransactionID,
			GrossAmt: int64(payment.Amount * 1.1),
		},
		CustomerDetail: &midtrans.CustomerDetails{
			FName: brand.PicName,
			Email: brand.Email,
			Phone: brand.PicPhone,
		},
	}

	// Kirim permintaan transaksi ke Midtrans
	snapResp, err := snapGateway.CreateTransaction(snapReq)
	if err != nil {
		log.Println("Midtrans error:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create transaction"})
		return
	}

	// Update status pembayaran ke pending
	config.DB.Model(&payment).Update("status", "pending")

	// Berikan respons sukses
	c.JSON(http.StatusOK, gin.H{
		"message":        "Payment initiated",
		"order_id":       payment.TransactionID,
		"transaction_id": snapResp.Token,
		"payment_url":    snapResp.RedirectURL,
	})
}

// WebhookPaymentHandler menangani notifikasi dari Midtrans
func WebhookPaymentHandler(c *gin.Context) {
	var notification map[string]interface{}

	// Parse request body dari Midtrans
	if err := c.ShouldBindJSON(&notification); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Ambil order_id dari notifikasi
	orderID, ok := notification["order_id"].(string)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order_id"})
		return
	}

	// Ambil status transaksi dari notifikasi
	transactionStatus, _ := notification["transaction_status"].(string)

	// Log notifikasi untuk debugging
	log.Printf("Received webhook: order_id=%s, status=%s\n", orderID, transactionStatus)

	// Cek status transaksi
	if transactionStatus == "settlement" || transactionStatus == "capture" {
		// Pembayaran berhasil, update status di database
		config.DB.Model(&models.Payment{}).Where("transaction_id = ?", orderID).Update("status", "success")
		log.Printf("Payment successful: order_id=%s\n", orderID)
	} else if transactionStatus == "pending" {
		config.DB.Model(&models.Payment{}).Where("transaction_id = ?", orderID).Update("status", "pending")
		log.Printf("Payment pending: order_id=%s\n", orderID)
	} else if transactionStatus == "expire" || transactionStatus == "cancel" || transactionStatus == "failure" {
		config.DB.Model(&models.Payment{}).Where("transaction_id = ?", orderID).Update("status", "failed")
		log.Printf("Payment failed: order_id=%s\n", orderID)
	}

	// Berikan respons sukses ke Midtrans
	c.JSON(http.StatusOK, gin.H{"message": "Webhook received successfully"})
}
