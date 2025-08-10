package db

import (
	"log"

	"fitnessapp/models"
	"fitnessapp/seeddata"
	"gorm.io/gorm"
)

// Migrate runs database migrations and seeds initial data
func Migrate(db *gorm.DB) error {
	log.Println("Running database migrations...")
	
	// Auto-migrate the User model
	if err := db.AutoMigrate(&models.User{}); err != nil {
		return err
	}
	
	log.Println("Database migrations completed successfully")
	
	// Seed initial data
	log.Println("Seeding initial data...")
	if err := seeddata.SeedUsers(db); err != nil {
		return err
	}
	
	log.Println("Database seeding completed successfully")
	return nil
}