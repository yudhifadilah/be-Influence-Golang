package controllers

import (
	"influencer-golang/config"
	"influencer-golang/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetAllFAQs - Mengambil semua data FAQ
func GetAllFAQs(c *gin.Context) {
	var faqs []models.FAQ
	if err := config.DB.Find(&faqs).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve FAQs"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"faqs": faqs})
}

// GetFAQByID - Mengambil data FAQ berdasarkan ID
func GetFAQByID(c *gin.Context) {
	id := c.Param("id")
	var faq models.FAQ
	if err := config.DB.First(&faq, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "FAQ not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"faq": faq})
}

// CreateFAQ - Menambahkan FAQ baru
func CreateFAQ(c *gin.Context) {
	var faq models.FAQ
	if err := c.ShouldBindJSON(&faq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Create(&faq).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create FAQ"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "FAQ created successfully", "faq": faq})
}

// UpdateFAQ - Memperbarui FAQ berdasarkan ID
func UpdateFAQ(c *gin.Context) {
	id := c.Param("id")
	var faq models.FAQ

	if err := config.DB.First(&faq, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "FAQ not found"})
		return
	}

	if err := c.ShouldBindJSON(&faq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Save(&faq)
	c.JSON(http.StatusOK, gin.H{"message": "FAQ updated successfully", "faq": faq})
}

// DeleteFAQ - Menghapus FAQ berdasarkan ID
func DeleteFAQ(c *gin.Context) {
	id := c.Param("id")
	var faq models.FAQ

	if err := config.DB.First(&faq, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "FAQ not found"})
		return
	}

	config.DB.Delete(&faq)
	c.JSON(http.StatusOK, gin.H{"message": "FAQ deleted successfully"})
}
