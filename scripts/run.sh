#!/bin/bash

# Check if .env file exists, if not, create it from example
if [ ! -f .env ]; then
    echo "Creating .env file from .env.example..."
    cp .env.example .env
    echo "Please edit .env file with your configuration."
fi

# Build and run the application
echo "Building and running the URL shortener service..."
go run cmd/server/main.go 
