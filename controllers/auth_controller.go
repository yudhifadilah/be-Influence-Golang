package controllers

import (
	"log"
	"net/http"

	"influencer-golang/config"
	"influencer-golang/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// Role default untuk pengguna baru
const DefaultUserRole = "admin"

func Register(c *gin.Context) {
	var user models.User

	// Bind JSON ke struct user
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
		return
	}

	// Cek apakah username sudah ada
	var existingUser models.User
	err := config.DB.Where("username = ?", user.Username).First(&existingUser).Error
	if err == nil {
		c.JSON(http.StatusConflict, gin.H{"message": "Username sudah digunakan"})
		return
	} else if err != nil && err.Error() != "record not found" {
		log.Println("Database error:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Gagal memeriksa username"})
		return
	}

	// Hash password sebelum menyimpan ke database
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("Hashing error:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Gagal mengenkripsi password"})
		return
	}
	user.Password = string(hashedPassword)
	user.Role = DefaultUserRole // Set default role sebagai admin

	// Simpan user ke database
	if err := config.DB.Create(&user).Error; err != nil {
		log.Println("Database insert error:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Gagal membuat user"})
		return
	}

	// Response sukses
	c.JSON(http.StatusCreated, gin.H{
		"message": "Registrasi berhasil",
		"user_id": user.ID,
		"role":    user.Role,
	})
}

// Login untuk autentikasi user
func Login(c *gin.Context) {
	var input struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	if err := config.DB.Where("username = ?", input.Username).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Username atau password salah"})
		return
	}

	// Bandingkan password yang diinput dengan hash di database
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Username atau password salah"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Login berhasil", "user_id": user.ID, "role": user.Role})
}

// GetAllUsers untuk mendapatkan daftar semua user (hanya bisa diakses oleh admin)
func GetAllUsers(c *gin.Context) {
	var users []models.User
	if err := config.DB.Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Gagal mengambil data user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"users": users})
}

// GetUserByID untuk mendapatkan user berdasarkan ID
func GetUserByID(c *gin.Context) {
	id := c.Param("id")
	var user models.User

	if err := config.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "User tidak ditemukan"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}

// UpdateUser untuk mengedit user berdasarkan ID
func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User

	if err := config.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "User tidak ditemukan"})
		return
	}

	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user.Username = input.Username
	user.Role = input.Role // Admin bisa mengubah role user lain

	if err := config.DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Gagal memperbarui user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User berhasil diperbarui", "user": user})
}

// DeleteUser untuk menghapus user berdasarkan ID
func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User

	if err := config.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "User tidak ditemukan"})
		return
	}

	if err := config.DB.Delete(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Gagal menghapus user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User berhasil dihapus"})
}
