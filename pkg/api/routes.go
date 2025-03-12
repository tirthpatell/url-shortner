package api

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

// SetupRoutes configures the API routes
func SetupRoutes(router *gin.Engine, database *sql.DB) {
	// API routes
	api := router.Group("/api")
	{
		// Public endpoints
		api.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"status": "ok",
			})
		})

		// Protected endpoints (require API key)
		protected := api.Group("")
		protected.Use(APIKeyAuthMiddleware())
		{
			// Shorten URL endpoint
			protected.POST("/shorten", ShortenURLHandler(database))

			// URL statistics endpoint
			protected.GET("/stats/:shortPath", URLStatsHandler(database))
		}
	}

	// Redirect endpoint with /s/ prefix
	router.GET("/s/:shortPath", RedirectHandler(database))

	// Serve static files if needed
	// router.Static("/static", "./static")

	// Handle 404
	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Page not found",
		})
	})
}
