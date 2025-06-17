package http

import (
	middleware "be/services/items/middlewares"

	"github.com/gin-gonic/gin"
)

func ItemRoutes(r *gin.Engine) {
	itemGroup := r.Group("/item")
	{
		itemGroup.GET("/getAllItems", GetAllItemsHandler)
		itemGroup.GET("/getItemById/:itemId", middleware.ValidateAccessToken(), middleware.ItemIDMiddleware(), GetItemByIdHandler)

		itemGroup.POST("/createNewItem", middleware.ValidateItemFields(), CreateItemHandler)

		itemGroup.PUT("/:id", middleware.ValidateAccessToken(), middleware.ValidateItemFields(), UpdateItemByIdHandler)
		itemGroup.DELETE("/:id", middleware.ValidateAccessToken(), middleware.ItemIDMiddleware(), DeleteItemHandler)
		itemGroup.POST("/import", middleware.ValidateAccessToken(), middleware.RequireExcelFileMiddleware(), UploadExcelHandler)
		itemGroup.GET("/getErrorItems", middleware.ValidateAccessToken(), GetImportErrorsHandler)

		itemGroup.GET("/getItemDesc", middleware.ValidateAccessToken(), GetItemDescHandler) // lấy 3 item có view cao nhất

		itemGroup.POST("/searchItemByName", SearchItemByNameHandler)
		itemGroup.POST("/filterItemsByPrice", FilterItemsByPriceHandler)

	}
}
