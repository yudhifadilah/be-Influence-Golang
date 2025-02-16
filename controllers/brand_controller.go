package controllers

import (
	"influencer-golang/config"
	"influencer-golang/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// Register Brand
func RegisterBrand(c *gin.Context) {
	var brand models.Brand

	// Pastikan menggunakan form-data untuk binding
	if err := c.Request.ParseMultipartForm(10 << 20); err != nil { // 10MB max
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse form data"})
		return
	}

	// Ambil data dari form secara manual
	brand.Email = c.PostForm("email")
	brand.BrandName = c.PostForm("brand_name")
	brand.PicName = c.PostForm("pic_name")
	brand.PicPhone = c.PostForm("pic_phone")
	brand.Province = c.PostForm("province")
	brand.City = c.PostForm("city")
	brand.Password = c.PostForm("password")

	// Validasi input
	if brand.Email == "" || brand.BrandName == "" || brand.PicName == "" || brand.PicPhone == "" || brand.Province == "" || brand.City == "" || brand.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "All fields are required"})
		return
	}

	// Hash Password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(brand.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}
	brand.Password = string(hashedPassword)

	// Tetapkan role secara otomatis
	brand.Role = "brand"

	// Handle file upload jika ada
	file, err := c.FormFile("brand_logo")
	if err == nil {
		// Path penyimpanan
		path := "uploads/brands/" + file.Filename
		if err := c.SaveUploadedFile(file, path); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload logo"})
			return
		}
		brand.BrandLogo = path
	}

	// Simpan ke database
	if err := config.DB.Create(&brand).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register brand", "details": err.Error()})
		return
	}

	// Response sukses
	c.JSON(http.StatusCreated, gin.H{
		"message":    "Brand registered successfully",
		"brand_logo": brand.BrandLogo,
		"data":       brand,
	})
}

// Login Brand
func LoginBrand(c *gin.Context) {
	var input models.Brand
	var brand models.Brand

	// Bind JSON ke struct LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Cari user berdasarkan email
	if err := config.DB.Where("email = ?", input.Email).First(&brand).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Cek password
	if err := bcrypt.CompareHashAndPassword([]byte(brand.Password), []byte(input.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Login berhasil", "brand_id": brand.ID, "role": brand.Role})
}

// Get All Brands
func GetAllBrands(c *gin.Context) {
	var brands []models.Brand
	config.DB.Find(&brands)
	c.JSON(http.StatusOK, brands)
}

// Get Brand by ID
func GetBrandByID(c *gin.Context) {
	id := c.Param("id")
	var brand models.Brand
	if err := config.DB.First(&brand, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Brand not found"})
		return
	}
	c.JSON(http.StatusOK, brand)
}

func UpdateBrand(c *gin.Context) {
	id := c.Param("id")
	var brand models.Brand

	// Cek apakah brand ada
	if err := config.DB.First(&brand, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Brand not found"})
		return
	}

	// Parse data dari form-data
	if err := c.Request.ParseMultipartForm(10 << 20); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse form data"})
		return
	}

	// Data yang akan diperbarui
	updateData := map[string]interface{}{}

	// Ambil data dari form
	if email := c.PostForm("email"); email != "" {
		updateData["email"] = email
	}
	if brandName := c.PostForm("brand_name"); brandName != "" {
		updateData["brand_name"] = brandName
	}
	if picName := c.PostForm("pic_name"); picName != "" {
		updateData["pic_name"] = picName
	}
	if picPhone := c.PostForm("pic_phone"); picPhone != "" {
		updateData["pic_phone"] = picPhone
	}
	if province := c.PostForm("province"); province != "" {
		updateData["province"] = province
	}
	if city := c.PostForm("city"); city != "" {
		updateData["city"] = city
	}

	// Hash password jika ada
	if password := c.PostForm("password"); password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
			return
		}
		updateData["password"] = string(hashedPassword)
	}

	// Upload logo jika ada
	file, err := c.FormFile("brand_logo")
	if err == nil {
		path := "uploads/brands/" + file.Filename
		if err := c.SaveUploadedFile(file, path); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload logo"})
			return
		}
		updateData["brand_logo"] = path
	}

	// Update database
	if err := config.DB.Model(&brand).Updates(updateData).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update brand", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Brand updated successfully", "updated_data": updateData})
}

func DeleteBrand(c *gin.Context) {
	id := c.Param("id")
	var brand models.Brand

	// Cek apakah brand ada
	if err := config.DB.First(&brand, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Brand not found"})
		return
	}

	// Hard delete
	if err := config.DB.Unscoped().Delete(&brand).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete brand", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Brand deleted successfully"})
}
