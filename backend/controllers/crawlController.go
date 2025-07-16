package controllers

import (
    "crawler-backend/crawler"
    "crawler-backend/models"
    "net/http"

    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
)

type CrawlRequest struct {
    URL    string `json:"url" binding:"required"`
    UserID uint   `json:"user_id"`
}

func CrawlURL(db *gorm.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        var req CrawlRequest

        if err := c.ShouldBindJSON(&req); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        // Perform crawl
        result, err := crawler.CrawlPage(req.URL)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }

        // Assign the user ID
        result.UserID = req.UserID

        // Save result to DB
        if err := db.Create(&result).Error; err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save crawl result"})
            return
        }

        c.JSON(http.StatusOK, gin.H{"message": "Crawl successful", "data": result})
    }
}

func GetResults(db *gorm.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        var results []models.CrawlResult

        if err := db.Order("created_at desc").Find(&results).Error; err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch crawl results"})
            return
        }

        c.JSON(http.StatusOK, gin.H{"data": results})
    }
}
