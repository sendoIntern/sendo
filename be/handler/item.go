package handler

import (
	"be/db"
	"be/entity"
	"fmt"

	"github.com/gin-gonic/gin"
)

func GetItemsHandler(c *gin.Context){
	var items []entity.Item
	result := db.DB.Find(&items)
	if result.Error != nil {
		c.JSON(400, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(200, items)
}

func GetItemByIdHandler(c *gin.Context){
	var item entity.Item
	id := c.Param("itemId")
	result := db.DB.First(&item, "id = ?", id)
	if result.Error != nil {
		c.JSON(400, gin.H{"error": result.Error.Error()})
		return
	}
	db.DB.Model(&item).Update("view", item.View + 1)
	c.JSON(200, item)
} 

func CreateItemHandler(c *gin.Context){
	fmt.Println("CreateItemHandler")
}

func DeleteItemHandler(c *gin.Context){
	fmt.Println("DeleteItemHandler")
}

func UpdateItemByIdHandler(c *gin.Context){
	fmt.Println("UpdateItemByIdHandler")
}
