#!/bin/bash

# Build and start the Docker containers
echo "Building and starting Docker containers..."
docker-compose up -d --build

# Check if containers are running
echo "Checking container status..."
docker-compose ps

echo "URL shortener service is running at http://localhost:8080"
echo "API endpoints:"
echo "  - POST http://localhost:8080/api/shorten"
echo "  - GET http://localhost:8080/api/stats/{shortPath}"
echo "  - GET http://localhost:8080/{shortPath} (redirect)" 
