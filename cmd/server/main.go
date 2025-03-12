package main

import (
	"fmt"
	"log"
	"os"

	"url-shortener/pkg/api"
	"url-shortener/pkg/db"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found, using environment variables")
	}

	// Initialize database connection
	database, err := db.InitDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer database.Close()

	// Create tables if they don't exist
	if err := db.CreateTables(database); err != nil {
		log.Fatalf("Failed to create tables: %v", err)
	}

	// Set up Gin router
	router := gin.Default()

	// Initialize API routes
	api.SetupRoutes(router, database)

	// Get port from environment or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Start the server
	fmt.Printf("Server running on port %s\n", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
