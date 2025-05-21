package handler

import (
	"be/db"
	"be/dto"
	"be/entity"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
	//limit input from request
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

	//only show info that's need to show
	resp := dto.ItemCreationResponse{
		ID:          item.ID,
		Name:        item.Name,
		Description: item.Description,
		Quantity:    item.Quantity,
		Price:       item.Price,
		Picture:     item.Picture,
		CreatedAt:   item.CreatedAt,
	}

	c.JSON(http.StatusOK, gin.H{"message": "Item created", "item": resp})
}

func DeleteItemHandler(c *gin.Context) {
	id := c.Param("id")

	// validate uuid
	uid, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid item ID"})
		return
	}

	result := db.DB.Delete(&entity.Item{}, uid)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete item"})
		return
	}

	// check is deleted?
	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":        "Item deleted successfully",
		"row(s) deleted": result.RowsAffected,
	})
}

func UpdateItemByIdHandler(c *gin.Context) {
	id := c.Param("id")

	uid, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid item ID"})
		return
	}

	var req dto.ItemUpdatingRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var item entity.Item
	result := db.DB.First(&item, "id = ?", uid)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}

	// update fields
	item.Name = req.Name
	item.Description = req.Description
	item.Quantity = req.Quantity
	item.Price = req.Price
	item.Picture = req.Picture

	// Lưu thay đổi vào DB
	if err := db.DB.Save(&item).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update item"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Item updated successfully", "item": item})
}
