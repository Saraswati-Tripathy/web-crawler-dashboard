package controllers

import (
    "net/http"
    "os"
    "time"

    "github.com/gin-gonic/gin"
    "github.com/golang-jwt/jwt/v4"
    "golang.org/x/crypto/bcrypt"
    "gorm.io/gorm"

    "crawler-backend/models"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

func Register(db *gorm.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        var input models.User
        if err := c.ShouldBindJSON(&input); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
        input.Password = string(hashedPassword)

        if err := db.Create(&input).Error; err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not register user"})
            return
        }

        c.JSON(http.StatusOK, gin.H{"message": "Registration successful"})
    }
}

func Login(db *gorm.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        var input models.User
        var user models.User

        if err := c.ShouldBindJSON(&input); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        if err := db.Where("email = ?", input.Email).First(&user).Error; err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
            return
        }

        if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password"})
            return
        }

        token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
            "user_id": user.ID,
            "exp":     time.Now().Add(time.Hour * 72).Unix(),
        })

        tokenString, _ := token.SignedString(jwtSecret)

        c.JSON(http.StatusOK, gin.H{"token": tokenString})
    }
}
