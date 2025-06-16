package http

import (
	middleware "be/services/items/middlewares"

	"github.com/gin-gonic/gin"
)

func ItemRoutes(r *gin.Engine) {
	itemGroup := r.Group("/item")
	{
		itemGroup.GET("/getAllItems", GetAllItemsHandler)
		itemGroup.GET("/getItemById/:itemId", middleware.ItemIDMiddleware(), GetItemByIdHandler)

		itemGroup.POST("/createNewItem", middleware.ValidateItemFields(), CreateItemHandler)

		itemGroup.PUT("/:id", middleware.ValidateItemFields(), UpdateItemByIdHandler)
		itemGroup.DELETE("/:id", middleware.ItemIDMiddleware(), DeleteItemHandler)
		itemGroup.POST("/import", middleware.RequireExcelFileMiddleware(), UploadExcelHandler)
		itemGroup.GET("/getErrorItems", GetImportErrorsHandler)

		itemGroup.GET("/getItemDesc", GetItemDescHandler) // lấy 3 item có view cao nhất
	}
}
