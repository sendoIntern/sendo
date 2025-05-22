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

	route := gin.Default()

	route.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173", "http://127.0.0.1:5173"}, // FE origin
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	db.New()
	defer db.Close()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	//login GG
	auth := route.Group("/auth")
	{
		google := auth.Group("/google")
		{
			google.GET("/login", handler.GoogleLoginHandler)
			google.GET("/callback", handler.GoogleCallbackHandler)
		}
	}

	//CRUD
	item := route.Group("/item")
	{
		item.GET("/getAllItems", handler.GetItemsHandler) // get full item
		item.PATCH("/:itemId", handler.GetItemByIdHandler)
		item.POST("/", handler.CreateItemHandler)
		item.DELETE("/:id", handler.DeleteItemHandler)
		item.PUT("/:id", handler.UpdateItemByIdHandler)
	}

	route.Run(":8080")
}
