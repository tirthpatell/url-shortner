package db

import (
	"database/sql"
	"time"
)

// URL represents a shortened URL in the database
type URL struct {
	ID           int       `json:"id"`
	OriginalURL  string    `json:"original_url"`
	ShortPath    string    `json:"short_path"`
	CreatedAt    time.Time `json:"created_at"`
	LastAccessed time.Time `json:"last_accessed,omitempty"`
	AccessCount  int       `json:"access_count"`
}

// StoreURL stores a new URL in the database
func StoreURL(db *sql.DB, originalURL, shortPath string) (*URL, error) {
	query := `
	INSERT INTO urls (original_url, short_path)
	VALUES ($1, $2)
	RETURNING id, original_url, short_path, created_at, last_accessed, access_count
	`

	var url URL
	var lastAccessed sql.NullTime

	err := db.QueryRow(query, originalURL, shortPath).Scan(
		&url.ID,
		&url.OriginalURL,
		&url.ShortPath,
		&url.CreatedAt,
		&lastAccessed,
		&url.AccessCount,
	)

	if lastAccessed.Valid {
		url.LastAccessed = lastAccessed.Time
	}

	return &url, err
}

// GetURLByShortPath retrieves a URL by its short path
func GetURLByShortPath(db *sql.DB, shortPath string) (*URL, error) {
	query := `
	SELECT id, original_url, short_path, created_at, last_accessed, access_count
	FROM urls
	WHERE short_path = $1
	`

	var url URL
	var lastAccessed sql.NullTime

	err := db.QueryRow(query, shortPath).Scan(
		&url.ID,
		&url.OriginalURL,
		&url.ShortPath,
		&url.CreatedAt,
		&lastAccessed,
		&url.AccessCount,
	)

	if lastAccessed.Valid {
		url.LastAccessed = lastAccessed.Time
	}

	return &url, err
}

// UpdateURLStats updates the access statistics for a URL
func UpdateURLStats(db *sql.DB, shortPath string) error {
	query := `
	UPDATE urls
	SET access_count = access_count + 1, last_accessed = CURRENT_TIMESTAMP
	WHERE short_path = $1
	`

	_, err := db.Exec(query, shortPath)
	return err
}

// CheckShortPathExists checks if a short path already exists
func CheckShortPathExists(db *sql.DB, shortPath string) (bool, error) {
	query := `
	SELECT EXISTS(SELECT 1 FROM urls WHERE short_path = $1)
	`

	var exists bool
	err := db.QueryRow(query, shortPath).Scan(&exists)
	return exists, err
}
