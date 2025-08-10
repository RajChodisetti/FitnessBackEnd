#!/bin/bash

set -e

CONTAINER_NAME="fitness-postgres"
PORT=5434

# Function to cleanup
cleanup() {
    echo "Cleaning up..."
    docker stop $CONTAINER_NAME 2>/dev/null || true
    docker rm $CONTAINER_NAME 2>/dev/null || true
    exit 0
}

# Set trap for cleanup on exit
trap cleanup SIGINT SIGTERM EXIT

# Clean up any existing container
docker stop $CONTAINER_NAME 2>/dev/null || true
docker rm $CONTAINER_NAME 2>/dev/null || true

echo "Starting PostgreSQL container on port $PORT..."

# Start PostgreSQL container with --rm for automatic cleanup
docker run --rm \
    --name $CONTAINER_NAME \
    -e POSTGRES_USER=postgres \
    -e POSTGRES_PASSWORD=passw \
    -e POSTGRES_DB=fitness \
    -p $PORT:5432 \
    postgres:15 &

# Get the container PID
CONTAINER_PID=$!

echo "Waiting for PostgreSQL to be ready..."

# Wait for database to be ready
max_attempts=30
attempts=0
while [ $attempts -lt $max_attempts ]; do
    if docker exec $CONTAINER_NAME pg_isready -U postgres -d fitness > /dev/null 2>&1; then
        echo "✅ PostgreSQL is ready on port $PORT!"
        break
    fi
    attempts=$((attempts + 1))
    echo "Waiting... ($attempts/$max_attempts)"
    sleep 2
done

if [ $attempts -eq $max_attempts ]; then
    echo "❌ PostgreSQL failed to become ready"
    exit 1
fi

echo "Database connection info:"
echo "  Host: localhost"
echo "  Port: $PORT"
echo "  Database: fitness"
echo "  User: postgres"
echo "  Password: passw"
echo ""
echo "You can now run: go run server/main.go"
echo "Press Ctrl+C to stop and cleanup"

# Wait for the container to finish
wait $CONTAINER_PID