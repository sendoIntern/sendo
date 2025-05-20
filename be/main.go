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

	route.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // FE origin
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

<<<<<<< HEAD
	//login GG
=======
>>>>>>> 364e6f452d499519a1bd781d49d714bafe42541e
	auth := route.Group("/auth")
	{
		google := auth.Group("/google")
		{
			google.GET("/login", handler.GoogleLoginHandler)
			google.GET("/callback", handler.GoogleCallbackHandler)
		}
	}
<<<<<<< HEAD

	//CRUD
	item := route.Group("/item")
	{
		item.GET("/", handler.GetItemsHandler) // get full item
		item.GET("/:id", handler.GetItemByIdHandler)
		item.POST("/create", handler.CreateItemHandler)
		item.DELETE("/:id", handler.DeleteItemHandler)
		item.PUT("/:id", handler.UpdateItemByIdHandler)
	}
	
=======
>>>>>>> 364e6f452d499519a1bd781d49d714bafe42541e
	route.Run(":8080")
}
