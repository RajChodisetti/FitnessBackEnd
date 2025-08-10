package models

// User represents a user of the fitness application.
type User struct {
	UserID   uint   `gorm:"primaryKey"`
	Email    string `gorm:"uniqueIndex;not null"`
	Password string `gorm:"not null"`
	UserType string `gorm:"not null"`
}
