package http

import (
	"be/pkg/db"
	"be/services/items/model/entity"
	"be/services/items/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
	// "github.com/google/uuid"
)

func GetAllItemsHandler(c *gin.Context) {
	items, err := usecase.GetAllItems()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot get items: " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"items": items,
	})
	
	
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

// func CreateItemHandler(c *gin.Context) {
// 	file, fileHeader, err := c.Request.FormFile("picture")
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Image is required"})
// 		return
// 	}
// 	defer file.Close()

// 	imageURL, err := utils.UploadToCloudinary(file, fileHeader)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot upload image"})
// 		return
// 	}

// 	// Nhận các field khác từ form-data
// 	var req request.ItemCreationRequest
// 	req.Name = c.PostForm("name")
// 	req.Description = c.PostForm("description")
// 	req.Quantity, _ = strconv.ParseInt(c.PostForm("quantity"), 10, 64)
// 	req.Price, _ = strconv.ParseFloat(c.PostForm("price"), 64)

// 	item := entity.Item{
// 		Name:        req.Name,
// 		Description: req.Description,
// 		Quantity:    req.Quantity,
// 		Price:       req.Price,
// 		Picture:     imageURL,
// 	}

// 	if err := db.DB.Create(&item).Error; err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot create item"})
// 		return
// 	}

// 	resp := response.ItemCreationResponse{
// 		ID:          item.ID,
// 		Name:        item.Name,
// 		Description: item.Description,
// 		Quantity:    item.Quantity,
// 		Price:       item.Price,
// 		Picture:     item.Picture,
// 		CreatedAt:   item.CreatedAt,
// 	}

// 	c.JSON(http.StatusOK, gin.H{"message": "Item created", "item": resp})
// }

// func DeleteItemHandler(c *gin.Context) {
// 	id := c.Param("id")

// 	// validate uuid
// 	uid, err := uuid.Parse(id)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid item ID"})
// 		return
// 	}

// 	result := db.DB.Delete(&entity.Item{}, uid)
// 	if result.Error != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete item"})
// 		return
// 	}

// 	// check is deleted?
// 	if result.RowsAffected == 0 {
// 		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{
// 		"message":        "Item deleted successfully",
// 		"row(s) deleted": result.RowsAffected,
// 	})
// }

// func UpdateItemByIdHandler(c *gin.Context) {
// 	id := c.Param("id")
// 	uid, err := uuid.Parse(id)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid item ID"})
// 		return
// 	}

// 	var item entity.Item
// 	result := db.DB.First(&item, "id = ?", uid)
// 	if result.Error != nil {
// 		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
// 		return
// 	}

// 	var req request.ItemUpdatingRequest
// 	req.Name = c.PostForm("name")
// 	req.Description = c.PostForm("description")
// 	req.Quantity, _ = strconv.ParseInt(c.PostForm("quantity"), 10, 64)
// 	req.Price, _ = strconv.ParseFloat(c.PostForm("price"), 64)

// 	file, fileHeader, err := c.Request.FormFile("picture")
// 	if err == nil { //exist new picture then upload
// 		defer file.Close()
// 		imageURL, err := utils.UploadToCloudinary(file, fileHeader)
// 		if err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot upload image"})
// 			return
// 		}
// 		item.Picture = imageURL
// 	}

// 	// Update other fields
// 	item.Name = req.Name
// 	item.Description = req.Description
// 	item.Quantity = req.Quantity
// 	item.Price = req.Price

// 	if err := db.DB.Save(&item).Error; err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update item"})
// 		return
// 	}

// 	resp := response.ItemUpdatingResponse{
// 		Name:        item.Name,
// 		Description: item.Description,
// 		Quantity:    item.Quantity,
// 		Price:       item.Price,
// 		Picture:     item.Picture,
// 		UpdatedAt:   item.UpdatedAt,
// 	}

// 	c.JSON(http.StatusOK, gin.H{"message": "Item updated successfully", "item": resp})
// }
