package shortener

import (
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"fmt"
	"os"
	"strconv"
	"strings"

	"url-shortener/pkg/db"
)

// Default length for short paths
const DefaultShortPathLength = 8

// GenerateShortPath generates a random short path
func GenerateShortPath(database *sql.DB) (string, error) {
	// Get URL length from environment or use default
	lengthStr := os.Getenv("URL_LENGTH")
	length := DefaultShortPathLength // Default length is now 8
	if lengthStr != "" {
		parsedLength, err := strconv.Atoi(lengthStr)
		if err == nil && parsedLength > 0 {
			length = parsedLength
		}
	}

	// Try to generate a unique short path
	maxAttempts := 5
	for i := 0; i < maxAttempts; i++ {
		shortPath, err := generateRandomString(length)
		if err != nil {
			return "", err
		}

		// Check if the short path already exists
		exists, err := db.CheckShortPathExists(database, shortPath)
		if err != nil {
			return "", err
		}

		if !exists {
			return shortPath, nil
		}
	}

	return "", fmt.Errorf("failed to generate a unique short path after %d attempts", maxAttempts)
}

// ShortenURL creates a shortened URL
func ShortenURL(database *sql.DB, originalURL string) (*db.URL, error) {
	// Generate a random short path
	shortPath, err := GenerateShortPath(database)
	if err != nil {
		return nil, err
	}

	// Store the URL in the database
	return db.StoreURL(database, originalURL, shortPath)
}

// GetFullShortURL returns the full shortened URL including domain and /s/ prefix
func GetFullShortURL(shortPath string) string {
	domain := os.Getenv("DOMAIN")
	if domain == "" {
		baseURL := os.Getenv("BASE_URL")
		if baseURL == "" {
			baseURL = "http://localhost:8080"
		}
		return fmt.Sprintf("%s/s/%s", baseURL, shortPath)
	}
	return fmt.Sprintf("https://%s/s/%s", domain, shortPath)
}

// generateRandomString generates a random alphanumeric string of the specified length
func generateRandomString(length int) (string, error) {
	// Calculate how many bytes we need
	byteLength := (length * 6) / 8
	if (length*6)%8 != 0 {
		byteLength++
	}

	// Generate random bytes
	bytes := make([]byte, byteLength)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}

	// Encode to base64
	encoded := base64.URLEncoding.EncodeToString(bytes)
	// Remove padding and non-alphanumeric characters
	encoded = strings.ReplaceAll(encoded, "=", "")
	encoded = strings.ReplaceAll(encoded, "-", "a")
	encoded = strings.ReplaceAll(encoded, "_", "A")

	// Truncate to desired length
	if len(encoded) > length {
		encoded = encoded[:length]
	}

	return encoded, nil
}
