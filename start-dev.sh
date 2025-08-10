#!/bin/bash

# Cleanup function
cleanup() {
    echo "Shutting down services..."
    docker-compose down --volumes --remove-orphans
    docker system prune -f 2>/dev/null || true
    exit 0
}

# Set trap to cleanup on script exit
trap cleanup SIGINT SIGTERM EXIT

# Remove any existing containers and volumes
docker-compose down --volumes --remove-orphans 2>/dev/null || true

# Start the database
echo "Starting database..."
docker-compose up -d db

# Wait for database to be ready
echo "Waiting for database to be ready..."
max_attempts=30
attempts=0

while [ $attempts -lt $max_attempts ]; do
    if docker-compose exec -T db pg_isready -U postgres -d fitness; then
        echo "Database is ready!"
        break
    fi
    attempts=$((attempts + 1))
    echo "Attempt $attempts/$max_attempts - waiting..."
    sleep 2
done

if [ $attempts -eq $max_attempts ]; then
    echo "Database failed to become ready"
    exit 1
fi

echo "Database started successfully on port 5433"
echo "You can now run the backend with: go run server/main.go"
echo "Press Ctrl+C to stop and cleanup"

# Keep script running to maintain cleanup trap
while true; do
    sleep 1
done