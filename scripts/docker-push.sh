#!/bin/bash

# Exit the script if any command fails
set -e

# Load environment variables from the .env file located in the scripts folder
SCRIPT_DIR="$(dirname "$0")"
if [ -f "$SCRIPT_DIR/.env" ]; then
  export $(grep -v '^#' "$SCRIPT_DIR/.env" | xargs)
else
  echo "Error: .env file not found in $SCRIPT_DIR."
  exit 1
fi

# Check if DOCKER_USER and DOCKER_REPO are set
if [ -z "$DOCKER_USER" ] || [ -z "$DOCKER_REPO" ]; then
  echo "Error: DOCKER_USER or DOCKER_REPO not set in .env file."
  exit 1
fi

# Define the full image name
IMAGE_NAME="$DOCKER_USER/$DOCKER_REPO"

# Get the current Git tag (or fallback to "latest" if no tag is found)
VERSION=$(git describe --tags --abbrev=0 2>/dev/null || echo "latest")

# Define the full image tag
IMAGE_TAG="$IMAGE_NAME:$VERSION"

# Navigate to the root directory where the Dockerfile is located
cd "$SCRIPT_DIR/.."

# Set up Docker Buildx for multi-architecture builds
echo "Setting up Docker Buildx for multi-architecture builds..."
docker buildx create --name multiarch-builder --use || true

# Build and push multi-architecture image
echo "Building and pushing multi-architecture Docker image with tag: $IMAGE_TAG"
docker buildx build --platform linux/amd64,linux/arm64 -t $IMAGE_TAG --push .

echo "Docker image $IMAGE_TAG has been successfully published for multiple architectures!" 
