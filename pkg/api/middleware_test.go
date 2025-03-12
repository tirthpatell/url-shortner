package api

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestAPIKeyAuthMiddleware(t *testing.T) {
	// Set up Gin in test mode
	gin.SetMode(gin.TestMode)

	// Test cases
	testCases := []struct {
		name           string
		envAPIKey      string
		requestAPIKey  string
		headerType     string
		expectedStatus int
	}{
		{
			name:           "No API key required",
			envAPIKey:      "",
			requestAPIKey:  "",
			headerType:     "",
			expectedStatus: http.StatusOK,
		},
		{
			name:           "Valid API key in X-API-Key header",
			envAPIKey:      "test-api-key",
			requestAPIKey:  "test-api-key",
			headerType:     "X-API-Key",
			expectedStatus: http.StatusOK,
		},
		{
			name:           "Valid API key in Authorization header",
			envAPIKey:      "test-api-key",
			requestAPIKey:  "test-api-key",
			headerType:     "Authorization",
			expectedStatus: http.StatusOK,
		},
		{
			name:           "Invalid API key",
			envAPIKey:      "test-api-key",
			requestAPIKey:  "wrong-api-key",
			headerType:     "X-API-Key",
			expectedStatus: http.StatusUnauthorized,
		},
		{
			name:           "Missing API key",
			envAPIKey:      "test-api-key",
			requestAPIKey:  "",
			headerType:     "",
			expectedStatus: http.StatusUnauthorized,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Set environment API key
			os.Setenv("API_KEY", tc.envAPIKey)
			defer os.Unsetenv("API_KEY")

			// Create a new router
			router := gin.New()

			// Add the middleware and a test handler
			router.Use(APIKeyAuthMiddleware())
			router.GET("/test", func(c *gin.Context) {
				c.Status(http.StatusOK)
			})

			// Create a test request
			req, _ := http.NewRequest("GET", "/test", nil)

			// Add the appropriate header
			if tc.headerType == "X-API-Key" && tc.requestAPIKey != "" {
				req.Header.Set("X-API-Key", tc.requestAPIKey)
			} else if tc.headerType == "Authorization" && tc.requestAPIKey != "" {
				req.Header.Set("Authorization", "Bearer "+tc.requestAPIKey)
			}

			// Perform the request
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			// Check the status code
			if w.Code != tc.expectedStatus {
				t.Errorf("Expected status %d, got %d", tc.expectedStatus, w.Code)
			}
		})
	}
}
