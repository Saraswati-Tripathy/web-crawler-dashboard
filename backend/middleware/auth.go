package middleware

import (
    "net/http"
    "os"
    "strings"

    "github.com/gin-gonic/gin"
    "github.com/golang-jwt/jwt/v4"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        authHeader := c.GetHeader("Authorization")
        if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
            return
        }

        tokenString := strings.TrimPrefix(authHeader, "Bearer ")

       token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
    return []byte(jwtSecret), nil
})
if err != nil {
    c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
    return
}

        claims, ok := token.Claims.(jwt.MapClaims)
if !ok || !token.Valid {
    c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
    c.Abort()
    return
}

userIDFloat, ok := claims["user_id"].(float64)
if !ok {
    c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid user ID in token"})
    c.Abort()
    return
}

userID := uint(userIDFloat)
c.Set("userID", userID)
    }
}
