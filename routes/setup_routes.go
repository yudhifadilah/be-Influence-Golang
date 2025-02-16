package routes

import (
	"influencer-golang/controllers"

	"github.com/gin-gonic/gin"
)

// SetupRoutes mengatur semua rute aplikasi
func SetupRoutes(r *gin.Engine) {

	// Login
	authRoutes := r.Group("/auth")
	{
		authRoutes.POST("/register", controllers.Register)
		authRoutes.POST("/login", controllers.Login)

		authRoutes.GET("/", controllers.GetAllUsers)
		authRoutes.GET("/:id", controllers.GetUserByID)
		authRoutes.PUT("/:id", controllers.UpdateUser)
		authRoutes.DELETE("/:id", controllers.DeleteUser)
	}

	// Brand Endpoint
	brandGroup := r.Group("/brands")
	{
		brandGroup.POST("/register", controllers.RegisterBrand)
		brandGroup.POST("/login", controllers.LoginBrand)
		brandGroup.GET("/", controllers.GetAllBrands)
		brandGroup.GET("/:id", controllers.GetBrandByID)
		brandGroup.PUT("/:id", controllers.UpdateBrand)
		brandGroup.DELETE("/:id", controllers.DeleteBrand)
	}

	// Influencer Endpoint
	influencerGroup := r.Group("/influencers")
	{
		influencerGroup.POST("/register", controllers.RegisterInfluencer)
		influencerGroup.POST("/login", controllers.LoginInfluencer)
		influencerGroup.GET("/", controllers.GetInfluencers)
		influencerGroup.GET("/:id", controllers.GetInfluencer)
		influencerGroup.PUT("/:id", controllers.UpdateInfluencer)
		influencerGroup.DELETE("/:id", controllers.DeleteInfluencer)
	}

	// campaigns Endpoint
	campaigns := r.Group("/campaigns")
	{
		campaigns.POST("/create", controllers.RegisterCampaign)        // Menambahkan campaign baru
		campaigns.GET("/", controllers.GetAllCampaigns)                // Menampilkan semua campaign
		campaigns.GET("/:id", controllers.GetCampaignByID)             // Menampilkan campaign berdasarkan ID
		campaigns.PUT("/:id/status", controllers.UpdateStatusCampaign) // Mengupdate status campaign
		campaigns.DELETE("/delete/:id", controllers.DeleteCampaign)    // Menghapus campaign berdasarkan ID
	}

	// Service Endpoint
	serviceRoutes := r.Group("/services")
	{
		serviceRoutes.POST("/", controllers.CreateService)
		serviceRoutes.GET("/", controllers.GetAllServices)
		serviceRoutes.GET("/:id", controllers.GetServiceByID)
		serviceRoutes.PUT("/:id", controllers.UpdateService)
		serviceRoutes.DELETE("/:id", controllers.DeleteService)
	}

	// Pembayaran Midtrans
	midRoutes := r.Group("/api")
	{
		midRoutes.POST("/payment", controllers.CreatePayment)

	}

	// Web hook untuk merubah status pada database ketika sudah berhasil payment
	hookRoutes := r.Group("/webhook")
	{
		hookRoutes.POST("/payment", controllers.WebhookPaymentHandler)
	}

	// Article Endpoint
	articleGroup := r.Group("/api/articles")
	{
		articleGroup.GET("/", controllers.GetAllArticles)
		articleGroup.GET("/:id", controllers.GetArticleByID)
		articleGroup.POST("/", controllers.CreateArticle)
		articleGroup.PUT("/:id", controllers.UpdateArticle)
		articleGroup.DELETE("/:id", controllers.DeleteArticle)
	}

	// FAQs Endpoint
	faqGroup := r.Group("/api/faqs")
	{
		faqGroup.GET("/", controllers.GetAllFAQs)
		faqGroup.GET("/:id", controllers.GetFAQByID)
		faqGroup.POST("/", controllers.CreateFAQ)
		faqGroup.PUT("/:id", controllers.UpdateFAQ)
		faqGroup.DELETE("/:id", controllers.DeleteFAQ)
	}
}
