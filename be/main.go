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
	db.New()
	defer db.Close()

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

	auth := route.Group("/auth")
	{
		google := auth.Group("/google")
		{
			google.GET("/login", handler.GoogleLoginHandler)
			google.GET("/callback", handler.GoogleCallbackHandler)
		}
	}
	route.Run(":8080")
}
