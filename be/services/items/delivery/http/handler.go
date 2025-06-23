package http

import (
	"be/pkg/db"
	"be/pkg/pagination"
	"be/services/items/model/entity"
	"be/services/items/model/request"
	"be/services/items/model/response"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"time"

	"be/pkg/cache"
	"be/pkg/rabbitmq"

	"be/services/items/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllItemsHandler(c *gin.Context) {
	search := c.DefaultQuery("search", "") // search string for name or description
	minPriceStr := c.DefaultQuery("minPrice", "0")
	maxPriceStr := c.DefaultQuery("maxPrice", "0")

	minPrice, _ := strconv.ParseFloat(minPriceStr, 64) // min price filter
	maxPrice, _ := strconv.ParseFloat(maxPriceStr, 64) // max price filter

	pageStr := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", "6")

	page, _ := strconv.Atoi(pageStr)   // page index
	limit, _ := strconv.Atoi(limitStr) // limit number for a page

	paging := pagination.Paging{
		Page:  page,
		Limit: limit,
	}

	// ------------------- REDIS CACHE -------------------
	cacheKey := fmt.Sprintf("items:search=%s:minPrice=%s:maxPrice=%s:page=%s:limit=%s",
		search, minPriceStr, maxPriceStr, pageStr, limitStr)

	// Check cache
	cached, err := cache.GetCache(cacheKey)
	if err == nil && cached != "" {
		var cachedResp response.APIResponse
		if err := json.Unmarshal([]byte(cached), &cachedResp); err == nil {
			c.JSON(http.StatusOK, cachedResp)
			return
		}
	}
	// ---------------------------------------------------

	items, err := usecase.GetAllItems(&paging, search, minPrice, maxPrice)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.APIResponse{
			Status:  "Fail",
			Message: "Cannot get items",
			Error:   err.Error(),
		})
		return
	}

	resp := response.APIResponse{
		Status:     "Success",
		Pagination: paging,
		Data:       items,
	}

	// Save to Redis
	data, _ := json.Marshal(resp)
	cache.SetCache(cacheKey, data, 10*time.Minute)

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

	cache.ClearCacheByKey("items:*") // invalidate cache

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
	var errs []error
	items, parseErrs := usecase.ParseExcel(file)
	if parseErrs != nil {
		errs = append(errs, parseErrs...)
	}
	if len(items) != 0 {
		for i := range items {
			item := items[i]
			if err := rabbitmq.Publish(item); err != nil {
				log.Printf(" Publish error: %v\n", err)
				errs = append(errs, errors.New("Publish item error:"+string(rune(i))+"__"+err.Error()))
			}
		}
	}

	cache.ClearCacheByKey("items:*") // invalidate cache

	c.JSON(http.StatusOK, response.APIResponse{
		Status:  "Success",
		Message: "Items imported to queue successfully",
		Data: map[string]interface{}{
			"errors": errs,
		},
	})
}

func GetImportErrorsHandler(c *gin.Context) {
	importErrs, err := usecase.GetErrorItems()
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.APIResponse{
			Status:  "Fail",
			Message: "Cannot get error items",
			Error:   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, response.APIResponse{
		Status:  "Success",
		Message: "Import errors retrieved successfully",
		Data:    importErrs,
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
