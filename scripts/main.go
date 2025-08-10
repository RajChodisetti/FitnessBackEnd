package main

import (
        "flag"
        "fmt"
        "log"
        "os"

        "gorm.io/driver/postgres"
        "gorm.io/gorm"

        "fitnessapp/models"
        "fitnessapp/seeddata"
)

// main provides a simple CLI for database migrations and seeding.
func main() {
	migrate := flag.Bool("migrate", false, "create database tables")
	seed := flag.Bool("seed", false, "load seed data")
	flag.Parse()

        host := getEnv("DB_HOST", "localhost")
        user := getEnv("DB_USER", "postgres")
        password := getEnv("DB_PASSWORD", "passw")
        dbname := getEnv("DB_NAME", "fitness")
       port := getEnv("DB_PORT", "5432")

        dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	if *migrate {
		if err := db.AutoMigrate(&models.User{}); err != nil {
			log.Fatalf("migration failed: %v", err)
		}
		fmt.Println("database migrated")
	}

	if *seed {
		if err := seeddata.SeedUsers(db); err != nil {
			log.Fatalf("seeding failed: %v", err)
		}
		fmt.Println("seed data loaded")
	}
}

func getEnv(key, fallback string) string {
        if v, ok := os.LookupEnv(key); ok && v != "" {
                return v
        }
        return fallback
}
