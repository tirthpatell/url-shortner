package shortener

import (
	"os"
	"testing"
)

func TestGenerateRandomString(t *testing.T) {
	// Test with different lengths
	lengths := []int{4, 6, 8, 10}

	for _, length := range lengths {
		result, err := generateRandomString(length)
		if err != nil {
			t.Errorf("generateRandomString(%d) returned an error: %v", length, err)
		}

		if len(result) != length {
			t.Errorf("generateRandomString(%d) returned string of length %d, expected %d", length, len(result), length)
		}
	}

	// Test multiple generations to ensure they're different
	results := make(map[string]bool)
	for i := 0; i < 10; i++ {
		result, _ := generateRandomString(6)
		if results[result] {
			t.Errorf("generateRandomString generated duplicate string: %s", result)
		}
		results[result] = true
	}
}

func TestGetFullShortURL(t *testing.T) {
	// Test with domain set
	os.Setenv("DOMAIN", "short.example.com")
	shortPath := "abc123"
	expected := "https://short.example.com/s/abc123"
	result := GetFullShortURL(shortPath)
	if result != expected {
		t.Errorf("GetFullShortURL(%s) = %s, expected %s", shortPath, result, expected)
	}

	// Test with base URL set
	os.Setenv("DOMAIN", "")
	os.Setenv("BASE_URL", "http://localhost:8080")
	expected = "http://localhost:8080/s/abc123"
	result = GetFullShortURL(shortPath)
	if result != expected {
		t.Errorf("GetFullShortURL(%s) = %s, expected %s", shortPath, result, expected)
	}

	// Test with neither set (default)
	os.Setenv("BASE_URL", "")
	expected = "http://localhost:8080/s/abc123"
	result = GetFullShortURL(shortPath)
	if result != expected {
		t.Errorf("GetFullShortURL(%s) = %s, expected %s", shortPath, result, expected)
	}
}
