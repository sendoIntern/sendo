package http

import (
	"github.com/gin-gonic/gin"
)

func Register(r *gin.Engine) {
	itemGroup := r.Group("/item")
	{
		itemGroup.GET("/getAllItems", GetItemsHandler)
		itemGroup.GET("/getItemById/:itemId", GetItemByIdHandler)
		itemGroup.POST("/createNewItem", CreateItemHandler)
		itemGroup.PUT("/:id", UpdateItemByIdHandler)
		itemGroup.DELETE("/:id", DeleteItemHandler)
	}
}
