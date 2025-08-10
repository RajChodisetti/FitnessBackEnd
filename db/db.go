package db

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// ConnectSQLite establishes a SQLite database connection for testing
func ConnectSQLite() (*gorm.DB, error) {
	dbPath := getEnv("DB_PATH", "fitness.db")
	return gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
}

// Connect establishes a database connection using postgres or sqlite settings.
func Connect() (*gorm.DB, error) {
	dbType := getEnv("DB_TYPE", "sqlite")
	
	if dbType == "sqlite" {
		return ConnectSQLite()
	}
	
	// PostgreSQL configuration
	host := getEnv("DB_HOST", "localhost")
	user := getEnv("DB_USER", "postgres")
	password := getEnv("DB_PASSWORD", "passw")
	dbname := getEnv("DB_NAME", "fitness")
	port := getEnv("DB_PORT", "5432")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, port)
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}

func getEnv(key, fallback string) string {
	if v, ok := os.LookupEnv(key); ok && v != "" {
		return v
	}
	return fallback
}