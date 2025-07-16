package controllers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"crawler-backend/models"
)

// CrawlURL handles crawling logic
func CrawlURL(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input struct {
			URL string `json:"url"`
		}

		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}

		//  Extract userID from context (middleware should set it)
		userIDVal, exists := c.Get("userID")
		if !exists {
			fmt.Println("userID not found in context")
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		userID, ok := userIDVal.(uint)
		if !ok || userID == 0 {
			fmt.Println("Invalid userID type or value:", userIDVal)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid user ID"})
			return
		}

		// Build result with status "queued"
		result := models.CrawlResult{
			URL:               input.URL,
			Title:             "", // Placeholder
			HTMLVersion:       "Unknown",
			InternalLinks:     0,
			ExternalLinks:     0,
			InaccessibleLinks: 0,
			HasLoginForm:      false,
			Status:            "queued",
			UserID:            userID,
			CreatedAt:         time.Now(),
			UpdatedAt:         time.Now(),
		}

		fmt.Printf("Saving crawl result for userID %d: %s\n", userID, input.URL)

		if err := db.Create(&result).Error; err != nil {
			fmt.Println("DB error:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save crawl result"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Crawl request submitted"})
	}
}

// GetResults returns all crawl results for the authenticated user
func GetResults(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		userIDVal, exists := c.Get("userID")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		userID, ok := userIDVal.(uint)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid user ID"})
			return
		}

		var results []models.CrawlResult
		if err := db.Where("user_id = ?", userID).Find(&results).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch results"})
			return
		}

		c.JSON(http.StatusOK, results)
	}
}
