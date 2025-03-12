package api

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

// APIKeyAuthMiddleware is a middleware that checks for a valid API key
func APIKeyAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get API key from environment
		expectedAPIKey := os.Getenv("API_KEY")
		if expectedAPIKey == "" {
			// If API_KEY is not set in environment, consider it disabled
			c.Next()
			return
		}

		// Get API key from request header
		apiKey := c.GetHeader("X-API-Key")
		if apiKey == "" {
			// Try to get from Authorization header (Bearer token)
			authHeader := c.GetHeader("Authorization")
			if strings.HasPrefix(authHeader, "Bearer ") {
				apiKey = strings.TrimPrefix(authHeader, "Bearer ")
			}
		}

		// Check if API key is valid
		if apiKey == "" || apiKey != expectedAPIKey {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid or missing API key",
			})
			c.Abort()
			return
		}

		// API key is valid, continue
		c.Next()
	}
}
