package middleware

import (
	"back/helpers"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Access denied, no authorization token provided"})
		c.Abort()
		return
	}

	// Split the header to get the token part
	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header must be Bearer token"})
		c.Abort()
		return
	}

	token := parts[1]
	email, userID, err := helpers.ValidateToken(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token: " + err.Error()})
		c.Abort()
		return
	}

	c.Set("email", email)
	c.Set("userID", userID)
	c.Next()
}
