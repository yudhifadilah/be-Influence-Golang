package controllers

import (
	"influencer-golang/config"
	"influencer-golang/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateService (Menambahkan layanan baru)
func CreateService(c *gin.Context) {
	var input models.Service

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Create(&input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal membuat layanan"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Layanan berhasil dibuat", "data": input})
}

// GetAllServices (Mendapatkan semua layanan)
func GetAllServices(c *gin.Context) {
	var services []models.Service

	if err := config.DB.Find(&services).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil layanan"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": services})
}

// GetServiceByID (Mendapatkan layanan berdasarkan ID)
func GetServiceByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID layanan tidak valid"})
		return
	}

	var service models.Service
	if err := config.DB.First(&service, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Layanan tidak ditemukan"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": service})
}

// UpdateService (Memperbarui layanan)
func UpdateService(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID layanan tidak valid"})
		return
	}

	var service models.Service
	if err := config.DB.First(&service, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Layanan tidak ditemukan"})
		return
	}

	var input models.Service
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Model(&service).Updates(input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal memperbarui layanan"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Layanan berhasil diperbarui", "data": service})
}

func DeleteService(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID layanan tidak valid"})
		return
	}

	var service models.Service
	if err := config.DB.First(&service, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Layanan tidak ditemukan"})
		return
	}

	// Hard delete
	if err := config.DB.Unscoped().Delete(&service).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete brand", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Layanan berhasil dihapus"})
}
