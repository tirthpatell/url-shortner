# Multi-Architecture Docker Support

This document explains how the URL shortener service supports multiple CPU architectures through Docker images.

## Overview

The URL shortener Docker image is built for multiple architectures:
- `linux/amd64` - For standard x86_64 servers
- `linux/arm64` - For ARM-based systems like Apple Silicon Macs

This allows the same Docker image to run on different types of hardware without compatibility issues.

## How It Works

The `scripts/docker-push.sh` script uses Docker Buildx to create multi-architecture images:

```bash
# Set up Docker Buildx for multi-architecture builds
docker buildx create --name multiarch-builder --use || true

# Build and push multi-architecture image
docker buildx build --platform linux/amd64,linux/arm64 -t $IMAGE_TAG --push .
```

## Building Multi-Architecture Images

To build and push a multi-architecture image:

1. Make sure you have Docker Buildx installed and set up
2. Configure your Docker Hub credentials in `scripts/.env`:
   ```
   DOCKER_USER=your-docker-username
   DOCKER_REPO=url-shortener
   ```
3. Run the Docker push script:
   ```
   ./scripts/docker-push.sh
   ```

## Using the Multi-Architecture Image

When you pull the image, Docker will automatically select the correct architecture for your system:

```bash
docker pull yourusername/url-shortener:latest
```

In your Docker Compose file, simply reference the image:

```yaml
services:
  app:
    image: yourusername/url-shortener:latest
    # ... other configuration
```

## Troubleshooting

If you encounter platform compatibility issues:

1. Verify that the image supports your architecture:
   ```
   docker inspect yourusername/url-shortener:latest
   ```

2. Check that Docker Buildx is properly configured:
   ```
   docker buildx ls
   ```

3. Ensure you're using the latest version of Docker with multi-architecture support 
