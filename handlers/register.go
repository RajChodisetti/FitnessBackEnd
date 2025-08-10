package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"fitnessapp/models"
)

type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	UserType string `json:"userType"`
}

// Register creates a new user.
func Register(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req RegisterRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
			return
		}
		u := models.User{Email: req.Email, Password: req.Password, UserType: req.UserType}
		if err := db.Create(&u).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "could not create user"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "user created"})
	}
}
