package middleware

import (
	"net/http"

	"influencer-golang/config"
	"influencer-golang/models"

	"github.com/gin-gonic/gin"
)

// AdminOnly middleware untuk membatasi akses hanya ke admin
func AdminOnly() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, _ := c.Get("user_id")

		var user models.User
		if err := config.DB.First(&user, userID).Error; err != nil || user.Role != "admin" {
			c.JSON(http.StatusForbidden, gin.H{"message": "Akses ditolak, hanya admin yang dapat mengakses"})
			c.Abort()
			return
		}

		c.Next()
	}

}

// Middleware untuk membatasi akses hanya untuk influencer
func InfluencerOnly() gin.HandlerFunc {
	return func(c *gin.Context) {
		userRole, exists := c.Get("role")
		if !exists || userRole != "influencer" {
			c.JSON(http.StatusForbidden, gin.H{"error": "Akses ditolak, hanya influencer yang dapat mengakses"})
			c.Abort()
			return
		}
		c.Next()
	}
}
