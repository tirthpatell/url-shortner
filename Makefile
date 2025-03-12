.PHONY: build run docker-build docker-run docker-stop docker-push init-db test clean generate-api-key

# Default target
all: build

# Build the application
build:
	go build -o bin/url-shortener ./cmd/server

# Run the application locally
run:
	./scripts/run.sh

# Build Docker image
docker-build:
	docker-compose build

# Run with Docker
docker-run:
	./scripts/docker-run.sh

# Stop Docker containers
docker-stop:
	docker-compose down

# Push Docker image to Docker Hub
docker-push:
	./scripts/docker-push.sh

# Initialize database
init-db:
	./scripts/init-db.sh

# Generate API key
generate-api-key:
	./scripts/generate-api-key.sh

# Run tests
test:
	go test -v ./...

# Clean build artifacts
clean:
	rm -rf bin/
	go clean

# Download dependencies
deps:
	go mod download

# Help
help:
	@echo "Available targets:"
	@echo "  build           - Build the application"
	@echo "  run             - Run the application locally"
	@echo "  docker-build    - Build Docker image"
	@echo "  docker-run      - Run with Docker"
	@echo "  docker-stop     - Stop Docker containers"
	@echo "  docker-push     - Push Docker image to Docker Hub"
	@echo "  init-db         - Initialize database"
	@echo "  generate-api-key - Generate a secure API key"
	@echo "  test            - Run tests"
	@echo "  clean           - Clean build artifacts"
	@echo "  deps            - Download dependencies"
	@echo "  help            - Show this help message" 
