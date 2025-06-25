package http

import (
	"be/pkg/db"
	"be/pkg/pagination"
	"be/pkg/rabbitmq"
	"be/services/items/model/entity"
	"be/services/items/model/request"
	"be/services/items/model/response"
	"encoding/json"
	"fmt"
	"time"

	"be/pkg/cache"

	"be/services/items/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllItemsHandler(c *gin.Context) {
	// ======= Parse query params =======
	search := c.DefaultQuery("search", "")
	minPriceStr := c.DefaultQuery("minPrice", "0")
	maxPriceStr := c.DefaultQuery("maxPrice", "0")
	pageStr := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", "6")
	statusStr := c.DefaultQuery("status", "")

	minPrice, _ := strconv.ParseFloat(minPriceStr, 64)
	maxPrice, _ := strconv.ParseFloat(maxPriceStr, 64)
	page, _ := strconv.Atoi(pageStr)
	limit, _ := strconv.Atoi(limitStr)

	paging := pagination.Paging{
		Page:  page,
		Limit: limit,
	}

	// ======= Cache key setup =======
	baseKey := fmt.Sprintf("items:search=%s:min=%s:max=%s:status=%s", search, minPriceStr, maxPriceStr, statusStr)
	cacheKey := fmt.Sprintf("%s:page=%s:limit=%s", baseKey, pageStr, limitStr)
	totalKey := fmt.Sprintf("%s:originalTotal", baseKey)

	// ======= Try get from cache =======
	if cached, err := cache.GetCache(cacheKey); err == nil && cached != "" {
		var cachedResp response.APIResponse
		if err := json.Unmarshal([]byte(cached), &cachedResp); err == nil {
			c.JSON(http.StatusOK, cachedResp)
			return
		}
	}

	originalTotalStr, totalItemErr := cache.GetCache(totalKey)
	originalTotal, _ := strconv.ParseInt(originalTotalStr, 10, 64)
	if totalItemErr != nil || originalTotalStr == "" {
		originalTotal = 0
	}

	// ======= Query database =======
	items, err := usecase.GetAllItems(&paging, search, minPrice, maxPrice, originalTotal, statusStr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.APIResponse{
			Status:  "Fail",
			Message: "Cannot get items",
			Error:   err.Error(),
		})
		return
	}

	// ======= Cache original total if not exists =======
	if totalItemErr != nil || originalTotalStr == "" {
		_ = cache.SetCache(totalKey, paging.Total, 5*time.Minute)
	}

	// ======= Response =======
	resp := response.APIResponse{
		Status:     "Success",
		Pagination: paging,
		Data:       items,
	}

	// Save full response to Redis cache
	if data, err := json.Marshal(resp); err == nil {
		_ = cache.SetCache(cacheKey, data, 15*time.Minute)
	}

	c.JSON(http.StatusOK, resp)
}

func GetItemByIdHandler(c *gin.Context) {
	database := db.GetDB()
	var item entity.Item
	id := c.Param("id")
	result := database.First(&item, "id = ?", id)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, response.APIResponse{
			Status:  "Fail",
			Message: "Cannot get item",
			Error:   result.Error.Error(),
		})
		return
	}
	database.Model(&item).Update("view", item.View+1)
	c.JSON(http.StatusOK, response.APIResponse{
		Status:  "Success",
		Message: "Item retrieved successfully",
		Data:    item,
	})
}

func CreateItemHandler(c *gin.Context) {
	file, fileHeader, err := c.Request.FormFile("picture")
	if err != nil {
		c.JSON(http.StatusBadRequest, response.APIResponse{
			Status:  "Fail",
			Message: "Image is required",
			Error:   err.Error(),
		})
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
		c.JSON(http.StatusInternalServerError, response.APIResponse{
			Status:  "Fail",
			Message: "Item Creation Error",
			Error:   err.Error(),
		})
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

	c.JSON(http.StatusOK, response.APIResponse{
		Status:  "Success",
		Message: "Item created successfully",
		Data:    resp,
	})
}

func ChangeStatusItemHandler(c *gin.Context) {
	id := c.Param("id")

	item, err := usecase.ChangeStatusItem(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.APIResponse{
			Status:  "Fail",
			Message: "Change Item Status Error",
			Error:   err.Error(),
		})
		return
	}

	cache.ClearCacheByKey("items:*") // invalidate cache

	c.JSON(http.StatusOK, response.APIResponse{
		Status:  "Success",
		Message: "Change Item Status Successfully",
		Data:    item,
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
		c.JSON(http.StatusInternalServerError, response.APIResponse{
			Status:  "Fail",
			Message: "Item Update Error",
			Error:   err.Error(),
		})
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

	cache.ClearCacheByKey("items:*") // invalidate cache

	c.JSON(http.StatusOK, response.APIResponse{
		Status:  "Success",
		Message: "Item updated successfully",
		Data:    resp,
	})
}

func ConfirmImportExcel(c *gin.Context) {

	var req struct {
		ItemsStr []request.ItemUploadRequest `json:"items"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.APIResponse{
			Status:  "Fail",
			Message: "Invalid items list",
			Error:   err.Error(),
		})
		return
	}

	// publish item to rabbitmq
	var errs []entity.ImportError
	for i, itemStr := range req.ItemsStr {

		quantity, _ := strconv.ParseInt(itemStr.Quantity, 10, 64)
		price, _ := strconv.ParseFloat(itemStr.Price, 64)
		view, _ := strconv.ParseInt(itemStr.View, 10, 64)
		recommend, _ := strconv.ParseInt(itemStr.Recommend, 10, 64)

		item := entity.Item{
			Name:        itemStr.Name,
			Description: itemStr.Description,
			Picture:     itemStr.Picture,
			Quantity:    quantity,
			Price:       price,
			View:        view,
			Recommend:   recommend,
		}

		if err := rabbitmq.Publish(item); err != nil {
			errs = append(errs, entity.ImportError{
				Description: fmt.Sprintf("PUBLISH ERROR AT ITEM(%d)_%s: %s", i, item.Name, err.Error()),
			})
		}
	}

	c.JSON(http.StatusOK, response.APIResponse{
		Status:  "Success",
		Message: "Items published successfully",
		Data:    errs,
	})
}

func UploadExcelHandler(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, response.APIResponse{
			Status:  "Fail",
			Message: "File is required",
			Error:   err.Error(),
		})
		return
	}
	items, err := usecase.ParseExcel(file)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.APIResponse{
			Status:  "Fail",
			Message: "Parse excel file error",
			Error:   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, response.APIResponse{
		Status:  "Success",
		Message: "Upload excel file successfully",
		Data:    items,
	})
}

func GetItemDescHandler(c *gin.Context) {
	itemDesc, err := usecase.GetItemDesc()
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.APIResponse{
			Status:  "Fail",
			Message: "Cannot get items desc",
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.APIResponse{
		Status:  "Success",
		Message: "Items desc retrieved successfully",
		Data:    itemDesc,
	})
}
