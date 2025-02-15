package routes

import (
	"influencer-golang/controllers"

	"github.com/gin-gonic/gin"
)

// SetupRoutes mengatur semua rute aplikasi
func SetupRoutes(r *gin.Engine) {

	authRoutes := r.Group("/auth")
	{
		authRoutes.POST("/register", controllers.Register)
		authRoutes.POST("/login", controllers.Login)

		// Hanya admin yang bisa mengelola user lain
		authRoutes.GET("/", controllers.GetAllUsers)
		authRoutes.GET("/:id", controllers.GetUserByID)
		authRoutes.PUT("/:id", controllers.UpdateUser)
		authRoutes.DELETE("/:id", controllers.DeleteUser)
	}

	brandGroup := r.Group("/brands")
	{
		brandGroup.POST("/register", controllers.RegisterBrand)
		brandGroup.POST("/login", controllers.LoginBrand)
		brandGroup.GET("/", controllers.GetAllBrands)
		brandGroup.GET("/:id", controllers.GetBrandByID)
		brandGroup.PUT("/:id", controllers.UpdateBrand)
		brandGroup.DELETE("/:id", controllers.DeleteBrand)
	}

	influencerGroup := r.Group("/influencers")
	{
		influencerGroup.POST("/register", controllers.RegisterInfluencer)
		influencerGroup.POST("/login", controllers.LoginInfluencer)
		influencerGroup.GET("/", controllers.GetInfluencers)
		influencerGroup.GET("/:id", controllers.GetInfluencer)
		influencerGroup.PUT("/:id", controllers.UpdateInfluencer)
		influencerGroup.DELETE("/:id", controllers.DeleteInfluencer)
	}

	campaigns := r.Group("/campaigns")
	{
		campaigns.POST("/create", controllers.RegisterCampaign)        // Menambahkan campaign baru
		campaigns.GET("/", controllers.GetAllCampaigns)                // Menampilkan semua campaign
		campaigns.GET("/:id", controllers.GetCampaignByID)             // Menampilkan campaign berdasarkan ID
		campaigns.PUT("/:id/status", controllers.UpdateStatusCampaign) // Mengupdate status campaign
		campaigns.DELETE("/delete/:id", controllers.DeleteCampaign)    // Menghapus campaign berdasarkan ID
	}

	serviceRoutes := r.Group("/services")
	{
		serviceRoutes.POST("/", controllers.CreateService)
		serviceRoutes.GET("/", controllers.GetAllServices)
		serviceRoutes.GET("/:id", controllers.GetServiceByID)
		serviceRoutes.PUT("/:id", controllers.UpdateService)
		serviceRoutes.DELETE("/:id", controllers.DeleteService)
	}

}
