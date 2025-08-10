package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"fitnessapp/handlers"
)

// New creates a gin.Engine with all API routes registered.
func New(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	r.Use(cors.Default())
	r.POST("/login", handlers.Login(db))
	r.POST("/register", handlers.Register(db))
	r.Static("/swagger", "./public/swagger")
	r.Static("/static", "./public")
	return r
}
