package routes

import (
	"be/services/handler"

	"github.com/gin-gonic/gin"
)

 func ItemRoutes(router *gin.Engine){
	api := router.Group("/item")
	{
		api.GET("/getAllItems", handler.GetItemsHandler) // get full item
		api.PATCH("/getItemById/:itemId", handler.GetItemByIdHandler)
		api.POST("/createNewItem", handler.CreateItemHandler)
		api.DELETE("/:id", handler.DeleteItemHandler)
		api.PUT("/:id", handler.UpdateItemByIdHandler)
	}
 }