package http

import (
	"github.com/gin-gonic/gin"
)

func ItemRoutes(r *gin.Engine) {
	itemGroup := r.Group("/item")
	{
		itemGroup.GET("/getAllItems", GetAllItemsHandler)
		itemGroup.GET("/getItemById/:itemId", GetItemByIdHandler)
		itemGroup.POST("/createNewItem", CreateItemHandler)
		itemGroup.PUT("/:id", UpdateItemByIdHandler)
		itemGroup.DELETE("/:id", DeleteItemHandler)
		itemGroup.POST("/import", UploadExcelHandler)
		itemGroup.GET("/getErrorItems", GetErrorItemsHandler)
	}
}
