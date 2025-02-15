package controllers

import (
	"influencer-golang/config"
	"influencer-golang/models"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
)

// RegisterCampaign (Menambahkan campaign baru)
func RegisterCampaign(c *gin.Context) {
	var input models.Campaign

	// Parsing multipart form untuk menangani file upload
	if err := c.Request.ParseMultipartForm(10 << 20); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse form data"})
		return
	}

	// Mengambil data dari form
	input.Name = c.PostForm("name")
	input.Category = c.PostForm("category")
	input.StartDate = c.PostForm("start_date")
	input.EndDate = c.PostForm("end_date")
	input.Status = "pending" // Set status otomatis menjadi "pending"

	// Konversi InfluencerID ke uint
	influencerID, err := strconv.ParseUint(c.PostForm("influencer_id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid influencer ID"})
		return
	}
	input.InfluencerID = uint(influencerID)

	// Handle Upload File PDF
	file, err := c.FormFile("pdf_file")
	if err == nil {
		uploadDir := "uploads/pdf/"
		if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create upload directory"})
			return
		}

		// Gunakan nama unik untuk file agar tidak ada konflik
		filename := strconv.FormatUint(influencerID, 10) + "_" + file.Filename
		filePath := filepath.Join(uploadDir, filename)
		if err := c.SaveUploadedFile(file, filePath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
			return
		}
		input.PDFFile = filePath
	}

	// Simpan data ke database
	if err := config.DB.Create(&input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register campaign"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Campaign registered successfully", "data": input})
}

// UpdateStatusCampaign (Mengupdate status campaign)
func UpdateStatusCampaign(c *gin.Context) {
	id := c.Param("id")
	var campaign models.Campaign

	// Cek apakah campaign ada di database
	if err := config.DB.First(&campaign, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Campaign not found"})
		return
	}

	// Ambil status baru dari request
	newStatus := c.PostForm("status")
	if newStatus == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Status is required"})
		return
	}

	// Update status campaign
	campaign.Status = newStatus
	if err := config.DB.Save(&campaign).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update campaign status"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Campaign status updated successfully", "data": campaign})
}

// Get All Campaigns
func GetAllCampaigns(c *gin.Context) {
	var campaigns []models.Campaign
	config.DB.Find(&campaigns)
	c.JSON(http.StatusOK, gin.H{"campaigns": campaigns})
}

// Get Campaign by ID
func GetCampaignByID(c *gin.Context) {
	id := c.Param("id")
	var campaign models.Campaign
	if err := config.DB.First(&campaign, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Campaign not found"})
		return
	}
	c.JSON(http.StatusOK, campaign)
}

// Update Campaign (Memperbarui data campaign)
func UpdateCampaign(c *gin.Context) {
	id := c.Param("id")
	var campaign models.Campaign

	if err := config.DB.First(&campaign, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Campaign not found"})
		return
	}

	// Bind data baru dari request
	var input models.Campaign
	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Handle Upload PDF baru jika ada
	file, err := c.FormFile("pdf_file")
	if err == nil {
		uploadDir := "uploads/pdf/"
		if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create upload directory"})
			return
		}

		// Path penyimpanan file baru
		filePath := filepath.Join(uploadDir, file.Filename)
		if err := c.SaveUploadedFile(file, filePath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
			return
		}

		// Hapus file lama jika ada
		if campaign.PDFFile != "" {
			_ = os.Remove(campaign.PDFFile)
		}

		// Update path file PDF baru
		input.PDFFile = filePath
	}

	// Update data campaign
	config.DB.Model(&campaign).Updates(input)

	c.JSON(http.StatusOK, gin.H{"message": "Campaign updated successfully", "data": campaign})
}

// Delete Campaign (Menghapus campaign berdasarkan ID)
func DeleteCampaign(c *gin.Context) {
	id := c.Param("id")
	var campaign models.Campaign

	if err := config.DB.First(&campaign, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Campaign not found"})
		return
	}

	// Hapus file PDF jika ada
	if campaign.PDFFile != "" {
		_ = os.Remove(campaign.PDFFile)
	}

	// Hapus dari database
	config.DB.Delete(&campaign)

	c.JSON(http.StatusOK, gin.H{"message": "Campaign deleted successfully"})
}

// Update Campaign Status (Memperbarui status campaign saja)
func UpdateCampaignStatus(c *gin.Context) {
	id := c.Param("id")
	var campaign models.Campaign

	if err := config.DB.First(&campaign, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Campaign not found"})
		return
	}

	// Ambil status baru dari request
	newStatus := c.PostForm("status")
	if newStatus == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Status is required"})
		return
	}

	// Update status di database
	config.DB.Model(&campaign).Update("status", newStatus)

	c.JSON(http.StatusOK, gin.H{"message": "Campaign status updated successfully", "status": newStatus})
}
