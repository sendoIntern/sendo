package handler

import (
	"be/db"
	"be/entity"
	"be/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ItemRecommendationHandler struct {
	cohereService *utils.CohereService
}

func NewItemRecommendationHandler() *ItemRecommendationHandler {
	return &ItemRecommendationHandler{
		cohereService: utils.NewCohereService(),
	}
}

type RecommendationRequest struct {
	Query string `json:"query" binding:"required"`
}

// GetRecommendationsByQuery - Gợi ý sản phẩm dựa trên câu query
func (h *ItemRecommendationHandler) GetRecommendationsByQuery(c *gin.Context) {
	var req RecommendationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	var items []entity.Item
	if err := db.DB.Find(&items).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch items from database"})
		return
	}

	recommendations, err := h.cohereService.GetItemRecommendations(c.Request.Context(), req.Query, items)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get recommendations"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"recommendations": recommendations,
	})
}

// GetSimilarItems - Gợi ý sản phẩm tương tự dựa trên ID sản phẩm
func (h *ItemRecommendationHandler) GetSimilarItems(c *gin.Context) {
	itemID := c.Param("itemId")
	id, err := uuid.Parse(itemID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid item ID"})
		return
	}

	// Lấy thông tin sản phẩm đang xem
	var currentItem entity.Item
	if err := db.DB.First(&currentItem, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}

	// Lấy tất cả sản phẩm khác (trừ sản phẩm đang xem)
	var otherItems []entity.Item
	if err := db.DB.Where("id != ?", id).Find(&otherItems).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch items from database"})
		return
	}

	// Tạo prompt dựa trên thông tin sản phẩm đang xem
	query := fmt.Sprintf("Find 5 Item that are most similar to: %s - %s (Price: %.2f)",
		currentItem.Name, currentItem.Description, currentItem.Price)

	recommendations, err := h.cohereService.GetItemRecommendations(c.Request.Context(), query, otherItems)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get recommendations", "message": err.Error()})
		return
	}

	// Giới hạn số lượng sản phẩm được gợi ý (ví dụ: 5 sản phẩm)
	// if len(recommendations) > 5 {
	// 	recommendations = recommendations[:5]
	// }

	c.JSON(http.StatusOK, gin.H{
		"current_item":  currentItem,
		"similar_items": recommendations,
	})
}
