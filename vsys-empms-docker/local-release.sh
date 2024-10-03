#!/bin/bash

# Exit on any error
set -e

# Usage message
usage() {
  echo "Usage: $0 [-b BUILD_SCRIPT] [-c COMPOSE_FILE] [-p IMAGE_PREFIX]"
  echo "  -b BUILD_SCRIPT  Path to the build script (default: ./build.sh)"
  echo "  -c COMPOSE_FILE  Path to the Docker Compose file (default: docker-compose.yml)"
  echo "  -p IMAGE_PREFIX  Prefix of Docker images to remove"
  exit 1
}

# Default values
BUILD_SCRIPT="./build.sh"
COMPOSE_FILE="docker-compose.yml"
IMAGE_PREFIX="vsys-empms"

# Parse command-line arguments
while getopts ":b:c:p:" opt; do
  case $opt in
    b)
      BUILD_SCRIPT=$OPTARG
      ;;
    c)
      COMPOSE_FILE=$OPTARG
      ;;
    p)
      IMAGE_PREFIX=$OPTARG
      ;;
    \?)
      echo "Invalid option: -$OPTARG" >&2
      usage
      ;;
    :)
      echo "Option -$OPTARG requires an argument." >&2
      usage
      ;;
  esac
done

# Check if IMAGE_PREFIX is set
if [ -z "$IMAGE_PREFIX" ]; then
  echo "Error: IMAGE_PREFIX is required."
  usage
fi

# Run the build script
echo "--------Running build script '$BUILD_SCRIPT'-----------"
"$BUILD_SCRIPT"

# Stop all running containers related to the project
echo "--------Stopping old containers-----------"
docker-compose -f "$COMPOSE_FILE" down -v
echo "--------Old containers stopped-----------"

# List and remove images with names starting with the specified prefix
echo "--------Removing images with prefix '${IMAGE_PREFIX}'-----------"
docker images --format "{{.Repository}}:{{.Tag}} {{.ID}}" | grep "^${IMAGE_PREFIX}" | awk '{print $2}' | xargs -r docker rmi
echo "--------Removed unused images-----------"

# Build images for containers
echo "--------Building images-----------"
docker-compose -f "$COMPOSE_FILE" build

# Start containers
echo "--------Starting containers-----------"
docker-compose -f "$COMPOSE_FILE" up

echo "--------Containers are up and running-----------"
