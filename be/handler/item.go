package handler

import (
	"be/db"
	"fmt"

	"github.com/gin-gonic/gin"
)

func GetItemsHandler(c *gin.Context){
	
	var items []items
	if err := db.Table("items").Find(&post).Error; err != nil {
		c.JSON(200, gin.H{
			"err": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{"data": items})
	
}

func GetItemByIdHandler(c *gin.Context){
	fmt.Println("GetItemByIdHandler")
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
