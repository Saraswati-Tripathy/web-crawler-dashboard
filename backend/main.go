package main

import (
	"crawler-backend/controllers"
	"crawler-backend/middleware"
	"crawler-backend/models"
	"github.com/gin-gonic/gin"
	
    "github.com/gin-contrib/cors"
    "time"
)

func main() {
	// Connect to the database
	models.ConnectDatabase()

	// Initialize Gin
	r := gin.Default()

	// Apply CORS middleware
	r.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"http://localhost:5173","http://127.0.0.1:5173"}, // your frontend origin
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
        AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
        MaxAge: 12 * time.Hour,
    }))

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
	r.Run(":8082") // Or any port you prefer
}
