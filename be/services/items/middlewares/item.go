package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"be/pkg/db"
	"be/services/items/model/entity"
	"errors"

	"github.com/google/uuid"
)

// function kiểm tra ID trước khi gọi handlers
func ItemIDMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method == "GET" && c.FullPath() == "/item/getItemById/:itemId" {
		
		id := c.Param("id")
		if id == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ID is required"})
			c.Abort()
			return
		}

		// Kiểm tra định dạng UUID
		if _, err := uuid.Parse(id); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
			c.Abort()
			return
		}
	
		// Kiểm tra xem item có tồn tại không
		var item entity.Item
		result := db.DB.First(&item, "id = ?", id)
		if result.Error != nil {
			if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error: " + result.Error.Error()})
			}
			c.Abort()
			return
		}
	}
		c.Next()
	}
}

// check input của item có hợp lệ không
func ValidateItemFields() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method == "POST" && c.FullPath() == "/item/createNewItem" {
			name := c.PostForm("name")
			price := c.PostForm("price")
			quantity := c.PostForm("quantity")
			description := c.PostForm("description")

			if name == "" {
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Name is required"})
				return
			}
			if price == "" {
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Price is required"})
				return
			}
			if quantity == "" {
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Quantity is required"})
				return
			}
			if description == "" {
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Description is required"})
				return
			}
		}

		c.Next()
	}
}
