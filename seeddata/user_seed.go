package seeddata

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"

	"fitnessapp/models"
	"gorm.io/gorm"
)

// SeedUsers loads users from seeddata/users.json and inserts them into the database.
func SeedUsers(db *gorm.DB) error {
	path := filepath.Join("seeddata", "users.json")
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	var users []models.User
	if err := json.Unmarshal(data, &users); err != nil {
		return err
	}

	for _, u := range users {
		var existing models.User
		err := db.Where("email = ?", u.Email).First(&existing).Error
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
		if errors.Is(err, gorm.ErrRecordNotFound) {
			if err := db.Create(&u).Error; err != nil {
				return err
			}
		}
	}
	return nil
}
