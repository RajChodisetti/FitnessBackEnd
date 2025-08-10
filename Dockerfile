# Stage 1: Build backend binary and migration tool
FROM golang:1.24.3-alpine AS builder
WORKDIR /app

# Dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy and build backend + migration tool
COPY . .

# Build the server and migration tool
RUN go build -o server ./server
RUN go build -o migrate ./scripts/main.go

# Stage 2: Runtime container - using golang image instead of alpine
FROM golang:1.24.3-alpine
WORKDIR /app

# Install required runtime libs
RUN apk add --no-cache bash postgresql-client

# Copy compiled binaries and assets
COPY --from=builder /app/server /app/server
COPY --from=builder /app/migrate /app/migrate
COPY --from=builder /app/seeddata /app/seeddata

# Make binaries executable and verify
RUN chmod 755 /app/server /app/migrate && \
    ls -la /app/server /app/migrate

# Set up environment variables
ENV DB_HOST=0.0.0.0
ENV DB_PORT=5432
ENV DB_USER=postgres
ENV DB_PASSWORD=passw
ENV DB_NAME=fitness

# Create a wait script
RUN echo '#!/bin/sh\n\
echo "Waiting for PostgreSQL..."\n\
sleep 30\n\
until pg_isready -h $DB_HOST -p $DB_PORT -U $DB_USER; do\n\
  echo "PostgreSQL not ready - sleeping"\n\
  sleep 5\n\
done\n\
echo "PostgreSQL ready - executing command"\n\
exec "$@"' > /wait-for-postgres.sh && chmod +x /wait-for-postgres.sh

EXPOSE 8080

CMD ["/bin/sh", "-c", "sleep 30 && /app/migrate -migrate -seed && /app/server"]