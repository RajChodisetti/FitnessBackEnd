# Fitness App Backend

A Go-based REST API backend for the Fitness App, built with Gin framework and PostgreSQL database.

## Tech Stack

- **Language**: Go 1.24.3
- **Framework**: Gin (HTTP web framework)
- **Database**: PostgreSQL
- **ORM**: GORM
- **Containerization**: Docker & Docker Compose

## Prerequisites

- Go 1.24.3+
- PostgreSQL 15+
- Docker & Docker Compose (for containerized setup)

## Installation & Setup

### Option 1: Docker Setup (Recommended)

1. Clone this repository:
```bash
git clone <repository-url>
cd fitness-app-backend
```

2. Run with Docker Compose:
```bash
docker-compose up --build
```

This will start:
- PostgreSQL database on port 5432
- Backend API server on port 8080

### Option 2: Local Development

1. Install dependencies:
```bash
go mod download
```

2. Start PostgreSQL database:
```bash
# Using Docker
docker run --name fitness-postgres \
  -e POSTGRES_USER=postgres \
  -e POSTGRES_PASSWORD=passw \
  -e POSTGRES_DB=fitness \
  -p 5432:5432 \
  -d postgres:15
```

3. Set environment variables:
```bash
export DB_HOST=localhost
export DB_PORT=5432
export DB_USER=postgres
export DB_PASSWORD=passw
export DB_NAME=fitness
```

4. Run database migrations and seed data:
```bash
go run ./scripts/main.go -migrate -seed
```

5. Start the server:
```bash
go run ./server/main.go
```

The API will be available at `http://localhost:8080`.

## API Endpoints

- `POST /api/register` - User registration
- `POST /api/login` - User authentication
- And more endpoints for fitness tracking features

## Project Structure

```
├── db/                     # Database connection and configuration
├── handlers/               # HTTP request handlers
├── models/                 # Database models (GORM)
├── router/                 # Route definitions
├── seeddata/              # Database seed data
├── server/                # Main server entry point
└── scripts/               # Database migration and seeding scripts
```

## Database

The application uses PostgreSQL with the following default configuration:
- Host: localhost (or 'db' in Docker)
- Port: 5432
- Username: postgres
- Password: passw
- Database: fitness

## Environment Variables

Set these environment variables for configuration:

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=passw
DB_NAME=fitness
```

## Development

For development with hot reload, you can use:
```bash
# Install air for hot reload (optional)
go install github.com/cosmtrek/air@latest

# Run with hot reload
air
```

## Testing

Run tests:
```bash
go test ./...
```

## Building

Build the application:
```bash
go build -o fitness-app-server ./server
./fitness-app-server
```