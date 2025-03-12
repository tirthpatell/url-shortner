#!/bin/bash

# This script generates a secure random API key and updates the .env file

# Generate a secure random API key (32 bytes, base64 encoded)
API_KEY=$(openssl rand -base64 32 | tr -d '\n')

echo "Generated API key: $API_KEY"

# Check if .env file exists
if [ -f .env ]; then
    # Check if API_KEY already exists in .env
    if grep -q "^API_KEY=" .env; then
        # Replace existing API_KEY
        sed -i.bak "s/^API_KEY=.*/API_KEY=$API_KEY/" .env
        rm -f .env.bak
        echo "Updated API_KEY in .env file"
    else
        # Add API_KEY to .env
        echo "API_KEY=$API_KEY" >> .env
        echo "Added API_KEY to .env file"
    fi
else
    # Create .env file from example
    if [ -f .env.example ]; then
        cp .env.example .env
        sed -i.bak "s/^API_KEY=.*/API_KEY=$API_KEY/" .env
        rm -f .env.bak
        echo "Created .env file from .env.example with new API_KEY"
    else
        # Create new .env file with just the API_KEY
        echo "API_KEY=$API_KEY" > .env
        echo "Created new .env file with API_KEY"
    fi
fi

echo "Your API key has been set. Use this key in your API requests."
echo "Example: curl -H \"X-API-Key: $API_KEY\" http://localhost:8080/api/shorten -d '{\"url\":\"https://example.com\"}' -H \"Content-Type: application/json\"" 
