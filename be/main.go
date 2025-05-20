package main

import (
	"be/db"
	"be/handler"
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	sql := db.New()
	sql.Connect()
	defer sql.Close()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	route := gin.Default()

	// Cấu hình CORS
	route.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // URL của frontend Vite
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Routes
	route.GET("/auth/google/login", handler.GoogleLoginHandler)
	route.GET("/auth/google/callback", handler.GoogleCallbackHandler)

	// Test route
	route.GET("/profile", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Welcome! You are logged in."})
	})

	route.Run(":8080")
}
