package main

import (
	"crawler-backend/controllers"
	"crawler-backend/middleware"
	"crawler-backend/models"
	"github.com/gin-gonic/gin"
)

func main() {
	// Connect to the database
	models.ConnectDatabase()

	// Initialize Gin
	r := gin.Default()

	// Apply CORS middleware
	r.Use(middleware.CORSMiddleware())

	// Public routes
	api := r.Group("/api")
	{
		api.POST("/register", controllers.Register(models.DB))
		api.POST("/login", controllers.Login(models.DB))
	}

	// Protected routes
	protected := r.Group("/api")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.POST("/crawl", controllers.CrawlURL(models.DB))
		protected.GET("/results", controllers.GetResults(models.DB))
	}

	// Start server
	r.Run(":8080") // Or any port you prefer
}
