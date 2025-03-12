# URL Shortener Service

A simple, efficient URL shortener service built with Go, PostgreSQL, and Docker.

> **Note:** This project was completely built with [Cursor](https://cursor.sh), an AI-powered code editor.

## Features

- Shorten long URLs to automatically generated 8-character alphanumeric codes
- URLs follow the format: https://sh.domain.com/s/{code}
- RESTful API for URL shortening and redirection
- API key authentication for protected endpoints
- Persistent storage with PostgreSQL
- Containerized with Docker for easy deployment
- Scalable architecture
- Built with Go 1.23 and latest dependencies
- Includes Postman collection for API testing

## Dependencies

This project uses the following key dependencies:

- Go 1.23.0 (with Go 1.24.1 toolchain)
- [gin-gonic/gin](https://github.com/gin-gonic/gin) v1.10.0 - HTTP web framework
- [lib/pq](https://github.com/lib/pq) v1.10.9 - PostgreSQL driver
- [joho/godotenv](https://github.com/joho/godotenv) v1.5.1 - .env file loader

All dependencies are kept up-to-date with the latest stable versions.

## Project Structure

```
url-shortener/
├── cmd/
│   └── server/           # Main application entry point
├── pkg/
│   ├── api/              # API handlers and routes
│   ├── db/               # Database connection and models
│   └── shortener/        # URL shortening logic
├── scripts/              # Utility scripts
├── postman/              # Postman collection for API testing
├── .env.example          # Example environment variables
├── docker-compose.yml    # Docker Compose configuration
├── Dockerfile            # Docker build configuration
├── go.mod                # Go module definition
└── README.md             # This file
```

## Prerequisites

- Go 1.23 or higher
- Docker and Docker Compose
- PostgreSQL (if running locally)
- Postman (for API testing)

## Getting Started

### Environment Setup

1. Clone the repository:
   ```
   git clone https://github.com/your-username/url-shortener.git
   cd url-shortener
   ```

2. Copy the example environment file and configure it:
   ```
   cp .env.example .env
   ```

3. Edit the `.env` file with your configuration, especially:
   - Set a secure `API_KEY` for authentication
   - Configure your database settings
   - Set your domain (default is sh.domain.com)

### Running with Docker

Build and start the service using Docker Compose:

```
docker-compose up -d
```

This will start both the URL shortener service and a PostgreSQL database.

### Publishing to Docker Hub

To build and publish the Docker image to Docker Hub:

1. Create a `.env` file in the `scripts` directory:
   ```
   cp scripts/.env.example scripts/.env
   ```

2. Edit the `scripts/.env` file with your Docker Hub credentials:
   ```
   DOCKER_USER=your-docker-username
   DOCKER_REPO=url-shortener
   ```

3. Run the Docker push script:
   ```
   ./scripts/docker-push.sh
   ```

This script will:
- Build a multi-architecture Docker image (supports both AMD64 and ARM64)
- Tag it with the latest Git tag (or "latest" if no tag exists)
- Push it to Docker Hub

For more details on multi-architecture support, see [MULTI_ARCH.md](MULTI_ARCH.md).

### Running Locally

1. Make sure PostgreSQL is running and accessible.

2. Install dependencies:
   ```
   go mod download
   ```

3. Run the server:
   ```
   go run cmd/server/main.go
   ```

## API Usage

### Authentication

Protected endpoints require an API key, which can be provided in one of two ways:

1. Using the `X-API-Key` header:
   ```
   X-API-Key: your_secure_api_key_here
   ```

2. Using the `Authorization` header with Bearer token:
   ```
   Authorization: Bearer your_secure_api_key_here
   ```

### Shorten a URL (Protected)

```
POST /api/shorten
Content-Type: application/json
X-API-Key: your_secure_api_key_here

{
  "url": "https://example.com/very/long/url/that/needs/shortening"
}
```

Response:
```
{
  "short_url": "https://sh.domain.com/s/a1b2c3d4",
  "original_url": "https://example.com/very/long/url/that/needs/shortening",
  "created_at": "2023-11-01T12:00:00Z"
}
```

### Get URL Statistics (Protected)

```
GET /api/stats/{shortPath}
X-API-Key: your_secure_api_key_here
```

### Get Original URL (Public)

```
GET /s/{shortPath}
```

This will redirect to the original URL. This endpoint is public and does not require authentication.

### Health Check (Public)

```
GET /api/health
```

Returns the status of the service. This endpoint is public and does not require authentication.

## Postman Collection

A Postman collection is included in the `postman` directory for easy testing of the API. The collection includes:

- All API endpoints with documentation
- Environment variables for easy configuration
- Example requests and responses
- Authentication setup for both API key and Bearer token methods

To use the collection:

1. Import the `postman/url-shortener-api.json` file into Postman
2. Set up your environment variables (see `postman/README.md` for details)
3. Start testing the API

For more details, see the [Postman Collection README](postman/README.md).

## Security Considerations

- Always use a strong, randomly generated API key
- Consider using HTTPS in production
- Regularly rotate your API keys
- Monitor for suspicious activity

## License

MIT 
