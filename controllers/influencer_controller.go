package controllers

import (
	"influencer-golang/config"
	"influencer-golang/models"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func RegisterInfluencer(c *gin.Context) {
	var input models.Influencer

	if err := c.Request.ParseMultipartForm(10 << 20); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse form data"})
		return
	}

	input.Email = c.PostForm("email")
	input.FullName = c.PostForm("full_name")
	input.Password = c.PostForm("password")
	input.BirthDate = c.PostForm("birth_date")
	input.Gender = c.PostForm("gender")
	input.InfluencerCategory = c.PostForm("influencer_category")
	input.PhoneNumber = c.PostForm("phone_number")
	input.KTPNumber = c.PostForm("ktp_number")
	input.NPWPNumber = c.PostForm("npwp_number")
	input.InstagramLink = c.PostForm("instagram_link")
	input.Role = "influencer" // Menambahkan role

	followersCount, err := strconv.Atoi(c.PostForm("followers_count"))
	if err == nil {
		input.FollowersCount = followersCount
	}

	input.BankAccount = c.PostForm("bank_account")
	input.AccountNumber = c.PostForm("account_number")
	input.Province = c.PostForm("province")
	input.City = c.PostForm("city")

	file, err := c.FormFile("profile_picture")
	if err == nil {
		dir := "uploads/"
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create upload directory"})
			return
		}

		filename := filepath.Join(dir, file.Filename)
		if err := c.SaveUploadedFile(file, filename); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
			return
		}
		input.ProfilePicture = filename
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}
	input.Password = string(hashedPassword)

	if err := config.DB.Create(&input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register influencer"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Influencer registered successfully"})
}

func LoginInfluencer(c *gin.Context) {
	var input models.Influencer
	var influencer models.Influencer

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := config.DB.Where("email = ?", input.Email).First(&influencer).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Email not found"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(influencer.Password), []byte(input.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"role":    influencer.Role,
	})
}

func GetInfluencers(c *gin.Context) {
	var influencers []models.Influencer
	config.DB.Find(&influencers)
	c.JSON(http.StatusOK, influencers)
}

func GetInfluencer(c *gin.Context) {
	id := c.Param("id")
	var influencer models.Influencer
	if err := config.DB.First(&influencer, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Influencer not found"})
		return
	}
	c.JSON(http.StatusOK, influencer)
}

func UpdateInfluencer(c *gin.Context) {
	id := c.Param("id")
	var influencer models.Influencer
	if err := config.DB.First(&influencer, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Influencer not found"})
		return
	}

	var input models.Influencer
	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updateData := map[string]interface{}{}
	if input.Email != "" {
		updateData["email"] = input.Email
	}
	if input.FullName != "" {
		updateData["full_name"] = input.FullName
	}
	if input.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
			return
		}
		updateData["password"] = string(hashedPassword)
	}

	file, err := c.FormFile("profile_picture")
	if err == nil {
		dir := "uploads/"
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create upload directory"})
			return
		}

		filename := filepath.Join(dir, file.Filename)
		if err := c.SaveUploadedFile(file, filename); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
			return
		}
		updateData["profile_picture"] = filename
	}

	config.DB.Model(&influencer).Updates(updateData)
	c.JSON(http.StatusOK, gin.H{"message": "Influencer updated successfully"})
}

func DeleteInfluencer(c *gin.Context) {
	id := c.Param("id")
	if err := config.DB.Delete(&models.Influencer{}, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Influencer not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Influencer deleted successfully"})
}
