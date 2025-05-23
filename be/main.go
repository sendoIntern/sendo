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
	//connect to db
	db.New()
	defer db.Close()

	//load env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	//create gin engine
	route := gin.Default()

	//cors
	route.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173", "http://127.0.0.1:5173"}, // FE origin
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	//login GG
	auth := route.Group("/auth")
	{
		google := auth.Group("/google")
		{
			google.GET("/login", handler.GoogleLoginHandler)
			google.GET("/callback", handler.GoogleCallbackHandler)
		}
	}

	//CRUD Item
	item := route.Group("/item")
	{
		item.GET("/getAllItems", handler.GetItemsHandler) // get full item
		item.PATCH("/:itemId", handler.GetItemByIdHandler)
		item.POST("/createNewItem", handler.CreateItemHandler)
		item.DELETE("/:id", handler.DeleteItemHandler)
		item.PUT("/:id", handler.UpdateItemByIdHandler)
	}

	// Item Recommendations
	recommendationHandler := handler.NewItemRecommendationHandler()
	recommendations := route.Group("/recommendations")
	{
		recommendations.POST("/search", recommendationHandler.GetRecommendationsByQuery) // Gợi ý dựa trên câu query
		recommendations.GET("/similar/:itemId", recommendationHandler.GetSimilarItems)   // Gợi ý sản phẩm tương tự
	}

	route.Run(":8080")
}
