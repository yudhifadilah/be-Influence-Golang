package controllers

import (
	"influencer-golang/config"
	"influencer-golang/models"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

// GetAllArticles mengambil semua artikel
func GetAllArticles(c *gin.Context) {
	var articles []models.Article
	if err := config.DB.Find(&articles).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve articles"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"articles": articles})
}

// GetArticleByID mengambil artikel berdasarkan ID
func GetArticleByID(c *gin.Context) {
	id := c.Param("id")
	var article models.Article
	if err := config.DB.First(&article, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Article not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"article": article})
}

// CreateArticle membuat artikel baru dengan upload gambar
func CreateArticle(c *gin.Context) {
	// Form input
	title := c.PostForm("title")
	excerpt := c.PostForm("excerpt")
	content := c.PostForm("content")

	// Handle file upload
	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Image file is required"})
		return
	}

	// Pastikan folder uploads/image/ ada
	imageDir := "uploads/image"
	if _, err := os.Stat(imageDir); os.IsNotExist(err) {
		os.MkdirAll(imageDir, os.ModePerm)
	}

	// Simpan file gambar dengan nama unik
	imagePath := filepath.Join(imageDir, file.Filename)
	if err := c.SaveUploadedFile(file, imagePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save image"})
		return
	}

	// Simpan artikel ke database
	article := models.Article{
		Title:   title,
		Excerpt: excerpt,
		Content: content,
		Image:   imagePath,
	}

	if err := config.DB.Create(&article).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create article"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Article created successfully", "article": article})
}

// UpdateArticle memperbarui artikel dengan upload gambar baru
func UpdateArticle(c *gin.Context) {
	id := c.Param("id")
	var article models.Article

	if err := config.DB.First(&article, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Article not found"})
		return
	}

	// Form input
	title := c.PostForm("title")
	excerpt := c.PostForm("excerpt")
	content := c.PostForm("content")

	// Cek apakah ada file gambar yang diunggah
	file, err := c.FormFile("image")
	if err == nil { // Jika ada gambar baru
		// Hapus gambar lama jika ada
		if article.Image != "" {
			os.Remove(article.Image)
		}

		// Simpan gambar baru
		imagePath := filepath.Join("uploads/image", file.Filename)
		if err := c.SaveUploadedFile(file, imagePath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save new image"})
			return
		}
		article.Image = imagePath
	}

	// Update data artikel
	article.Title = title
	article.Excerpt = excerpt
	article.Content = content

	config.DB.Save(&article)
	c.JSON(http.StatusOK, gin.H{"message": "Article updated successfully", "article": article})
}

// DeleteArticle menghapus artikel dan gambar terkait
func DeleteArticle(c *gin.Context) {
	id := c.Param("id")
	var article models.Article

	if err := config.DB.First(&article, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Article not found"})
		return
	}

	// Hapus gambar dari penyimpanan
	if article.Image != "" {
		os.Remove(article.Image)
	}

	// Hapus artikel dari database
	config.DB.Delete(&article)
	c.JSON(http.StatusOK, gin.H{"message": "Article deleted successfully"})
}
