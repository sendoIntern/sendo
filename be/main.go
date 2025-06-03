package main

import (
	"be/pkg/db"
	"be/services/items/delivery/http"
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173", "http://127.0.0.1:5173"}, // FE origin
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
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

	http.Register(router)

	// //login GG
	// auth := route.Group("/auth")
	// {
	// 	google := auth.Group("/google")
	// 	{
	// 		google.GET("/login", handler.GoogleLoginHandler)
	// 		google.GET("/callback", handler.GoogleCallbackHandler)
	// 	}
	// }
	// Login GG from frontend

	// routes.AuthRoutes(router)

	// auth := route.Group("/auth")
	// {
	// 	auth.POST("/login", handler.LoginHandler)
	// }
	//CRUD
	// item := route.Group("/item")
	// {
	// 	item.GET("/getAllItems", handler.GetItemsHandler) // get full item
	// 	item.PATCH("/getItemById/:itemId", handler.GetItemByIdHandler)
	// 	item.POST("/createNewItem", handler.CreateItemHandler)
	// 	item.DELETE("/:id", handler.DeleteItemHandler)
	// 	item.PUT("/:id", handler.UpdateItemByIdHandler)
	// }

	router.Run(":8080")
}
