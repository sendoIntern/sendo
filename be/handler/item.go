package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetItemsHandler(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{"message": "GetItemsHandler"})
	
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