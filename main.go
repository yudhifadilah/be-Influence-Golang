package main

import (
	"influencer-golang/config"
	"influencer-golang/routes"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("❌ Error loading .env file")
	}

	// Connect to the Database
	if err := config.ConnectDB(); err != nil {
		log.Fatalf("❌ Gagal terhubung ke database: %v", err)
	}

	log.Println("✅ Database terhubung dengan sukses!")

	// Inisialisasi Midtrans
	config.InitMidtrans()

	// Setup Router dengan CORS
	r := gin.Default()

	// Konfigurasi CORS agar lebih fleksibel
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // Ubah jika perlu batasan
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	// Setup Routes
	routes.SetupRoutes(r)

	// Ambil port dari environment atau gunakan default 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Jalankan server
	log.Printf("🚀 Server berjalan di port %s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("❌ Gagal menjalankan server: %v", err)
	}
}
