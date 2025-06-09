package http

import (
	"be/pkg/db"
	"be/services/items/model/entity"
	"be/services/items/model/request"
	"be/services/items/model/response"
	"log"

	"be/pkg/rabbitmq"

	"be/services/items/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllItemsHandler(c *gin.Context) {
	items, err := usecase.GetAllItems()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot get items: " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, items)

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
	file, fileHeader, err := c.Request.FormFile("picture")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Image is required"})
		return
	}
	defer file.Close()

	// Nhận các field khác từ form-data
	var req request.ItemCreationRequest
	req.Name = c.PostForm("name")
	req.Description = c.PostForm("description")
	req.Quantity, _ = strconv.ParseInt(c.PostForm("quantity"), 10, 64)
	req.Price, _ = strconv.ParseFloat(c.PostForm("price"), 64)
	req.PictureHeader = fileHeader
	req.PictureFile = &file

	item, err := usecase.CreateItem(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Item Creation Error": err.Error()})
		return
	}

	resp := response.ItemCreationResponse{
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

	err := usecase.DeleteItem(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Item Deletion Error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Item deleted successfully",
	})
}

func UpdateItemByIdHandler(c *gin.Context) {
	id := c.Param("id")

	var req request.ItemUpdatingRequest
	req.Name = c.PostForm("name")
	req.Description = c.PostForm("description")
	req.Quantity, _ = strconv.ParseInt(c.PostForm("quantity"), 10, 64)
	req.Price, _ = strconv.ParseFloat(c.PostForm("price"), 64)

	file, fileHeader, err := c.Request.FormFile("picture")
	if err == nil { //exist new picture then upload
		defer file.Close()
		req.PictureFile = &file
		req.PictureHeader = fileHeader
	} else {
		req.PictureFile = nil
		req.PictureHeader = nil
	}

	var item entity.Item
	item, err = usecase.UpdateItem(id, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Item Update Error": err.Error()})
		return
	}

	resp := response.ItemUpdatingResponse{
		Name:        item.Name,
		Description: item.Description,
		Quantity:    item.Quantity,
		Price:       item.Price,
		Picture:     item.Picture,
		UpdatedAt:   item.UpdatedAt,
	}

	c.JSON(http.StatusOK, gin.H{"message": "Item updated successfully", "item": resp})
}

func UploadExcelHandler(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "file is required"})
		return
	}
	items, err := usecase.ParseExcel(file)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}

	for i := range items {
		item := items[i]
		if err := rabbitmq.Publish(item); err != nil {
			log.Printf(" Publish error: %v\n", err)
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "imported to queue"})
}
