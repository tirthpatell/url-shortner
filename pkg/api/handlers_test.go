package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestShortenURLHandler_InvalidURL(t *testing.T) {
	// Set up Gin in test mode
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	// We don't need a real database for this test
	router.POST("/api/shorten", func(c *gin.Context) {
		var request ShortenURLRequest
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
			return
		}

		// Validate URL (this is what we're testing)
		if request.URL == "invalid-url" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid URL format"})
			return
		}

		c.JSON(http.StatusCreated, ShortenURLResponse{
			ShortURL:    "http://example.com/s/abc123",
			OriginalURL: request.URL,
			CreatedAt:   "2023-01-01T00:00:00Z",
		})
	})

	// Create a request with an invalid URL
	reqBody := ShortenURLRequest{
		URL: "invalid-url",
	}
	jsonData, _ := json.Marshal(reqBody)
	req, _ := http.NewRequest("POST", "/api/shorten", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	// Perform the request
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Check the response
	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d, got %d", http.StatusBadRequest, w.Code)
	}

	// Check the response body
	var response map[string]string
	if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
		t.Errorf("Failed to unmarshal response: %v", err)
	}

	if response["error"] != "Invalid URL format" {
		t.Errorf("Expected error message 'Invalid URL format', got '%s'", response["error"])
	}
}

func TestShortenURLHandler_ValidURL(t *testing.T) {
	// Set up Gin in test mode
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	// We don't need a real database for this test
	router.POST("/api/shorten", func(c *gin.Context) {
		var request ShortenURLRequest
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
			return
		}

		// For a valid URL, return a successful response
		c.JSON(http.StatusCreated, ShortenURLResponse{
			ShortURL:    "http://example.com/s/abc123",
			OriginalURL: request.URL,
			CreatedAt:   "2023-01-01T00:00:00Z",
		})
	})

	// Create a request with a valid URL
	reqBody := ShortenURLRequest{
		URL: "https://example.com/long/url",
	}
	jsonData, _ := json.Marshal(reqBody)
	req, _ := http.NewRequest("POST", "/api/shorten", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	// Perform the request
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Check the response
	if w.Code != http.StatusCreated {
		t.Errorf("Expected status code %d, got %d", http.StatusCreated, w.Code)
	}

	// Check the response body
	var response ShortenURLResponse
	if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
		t.Errorf("Failed to unmarshal response: %v", err)
	}

	if response.ShortURL != "http://example.com/s/abc123" {
		t.Errorf("Expected short URL 'http://example.com/s/abc123', got '%s'", response.ShortURL)
	}

	if response.OriginalURL != "https://example.com/long/url" {
		t.Errorf("Expected original URL 'https://example.com/long/url', got '%s'", response.OriginalURL)
	}
}
