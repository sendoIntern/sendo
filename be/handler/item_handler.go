package handler

import (
	"be/db"
	"be/dto"
	"be/entity"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetItemsHandler(c *gin.Context) {
	var items []entity.Item
	result := db.DB.Find(&items)
	if result.Error != nil {
		c.JSON(400, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(200, items)
}

func GetItemByIdHandler(c *gin.Context) {
	var item entity.Item
	id := c.Param("itemId")
	result := db.DB.First(&item, "id = ?", id)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
		return
	}
	db.DB.Model(&item).Update("view", item.View+1)
	c.JSON(http.StatusOK, item)
}

func CreateItemHandler(c *gin.Context) {
	var req dto.ItemCreationRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	item := entity.Item{
		Name:        req.Name,
		Description: req.Description,
		Quantity:    req.Quantity,
		Price:       req.Price,
		Picture:     req.Picture,
	}

	if err := db.DB.Create(&item).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot create item"})
		return
	}

	resp := dto.ItemCreationResponse{
		ID:          item.ID,
		Name:        item.Name,
		Description: item.Description,
		Quantity:    item.Quantity,
		Price:       item.Price,
		Picture:     item.Picture,
		CreateAt:    item.CreatedAt,
	}

	c.JSON(http.StatusOK, gin.H{"message": "Item created", "item": resp})
}

func DeleteItemHandler(c *gin.Context) {
	fmt.Println("DeleteItemHandler")
}

func UpdateItemByIdHandler(c *gin.Context) {
	fmt.Println("UpdateItemByIdHandler")
}
