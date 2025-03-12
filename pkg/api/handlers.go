package api

import (
	"database/sql"
	"net/http"
	"net/url"

	"url-shortener/pkg/db"
	"url-shortener/pkg/shortener"

	"github.com/gin-gonic/gin"
)

// ShortenURLRequest represents the request body for shortening a URL
type ShortenURLRequest struct {
	URL string `json:"url" binding:"required"`
}

// ShortenURLResponse represents the response for a shortened URL
type ShortenURLResponse struct {
	ShortURL    string `json:"short_url"`
	OriginalURL string `json:"original_url"`
	CreatedAt   string `json:"created_at"`
}

// ShortenURLHandler handles the URL shortening request
func ShortenURLHandler(database *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var request ShortenURLRequest
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
			return
		}

		// Validate URL
		_, err := url.ParseRequestURI(request.URL)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid URL format"})
			return
		}

		// Shorten URL
		urlObj, err := shortener.ShortenURL(database, request.URL)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Generate full short URL
		fullShortURL := shortener.GetFullShortURL(urlObj.ShortPath)

		// Return response
		c.JSON(http.StatusCreated, ShortenURLResponse{
			ShortURL:    fullShortURL,
			OriginalURL: urlObj.OriginalURL,
			CreatedAt:   urlObj.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
		})
	}
}

// RedirectHandler handles the URL redirection
func RedirectHandler(database *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		shortPath := c.Param("shortPath")

		// Get URL from database
		urlObj, err := db.GetURLByShortPath(database, shortPath)
		if err != nil {
			if err == sql.ErrNoRows {
				c.JSON(http.StatusNotFound, gin.H{"error": "Short URL not found"})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve URL"})
			}
			return
		}

		// Update access statistics asynchronously
		go func() {
			_ = db.UpdateURLStats(database, shortPath)
		}()

		// Redirect to original URL
		c.Redirect(http.StatusMovedPermanently, urlObj.OriginalURL)
	}
}

// URLStatsHandler returns statistics for a shortened URL
func URLStatsHandler(database *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		shortPath := c.Param("shortPath")

		// Get URL from database
		urlObj, err := db.GetURLByShortPath(database, shortPath)
		if err != nil {
			if err == sql.ErrNoRows {
				c.JSON(http.StatusNotFound, gin.H{"error": "Short URL not found"})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve URL"})
			}
			return
		}

		// Return URL statistics
		c.JSON(http.StatusOK, urlObj)
	}
}
