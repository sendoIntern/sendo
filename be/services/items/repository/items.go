package repository

import (
	"be/pkg/db"
	"be/pkg/pagination"
	"be/services/items/model/entity"
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

var database = db.GetDB()

func CreateItem(item entity.Item) error {
	if err := database.Create(&item).Error; err != nil {
		return err
	}
	return nil
}

func DeleteItem(id uuid.UUID) error {
	item, exist := IsExistItem(id)
	if !exist {
		return errors.New("ITEM NOT FOUND")
	}
	item.Status = false
	if err := database.Save(&item).Error; err != nil {
		return err
	}
	return nil
}

func ChangeStatusItem(id uuid.UUID) (entity.Item, error) {
	item, exist := IsExistItem(id)
	if !exist {
		return item, errors.New("ITEM NOT FOUND")
	}
	item.Status = !item.Status
	if err := database.Save(&item).Error; err != nil {
		return item, err
	}
	return item, nil
}

func UpdateItem(item entity.Item) error {
	if err := database.Save(&item).Error; err != nil {
		return err
	}
	return nil
}

func IsExistItem(id uuid.UUID) (entity.Item, bool) {
	var item entity.Item
	result := database.First(&item, "id = ?", id)
	if result.Error != nil {
		return entity.Item{}, false
	}
	return item, true
}

func IsExistItemByName(name string) bool {
	var item entity.Item
	result := database.First(&item, "name = ?", name)
	return result.Error == nil
}

// lấy 3 item có view cao nhất
func GetTopViewedItems() ([]entity.Item, error) {
	var items []entity.Item
	if err := database.Order("view DESC").Limit(3).Find(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}

// get items with search, filter and pagination
func FetchItems(
	p *pagination.Paging,
	search string,
	minPrice float64,
	maxPrice float64,
	originalTotal int64,
	status string) ([]entity.Item, error) {

	var items []entity.Item
	query := database.Model(&entity.Item{})
	if status != "" {
		query = query.Where("status = ?", status)
	}

	if search != "" {
		query = query.Where("name ILIKE ? OR description ILIKE ?", "%"+search+"%", "%"+search+"%")
	}

	if minPrice > 0 {
		query = query.Where("price >= ?", minPrice)
	}

	if maxPrice > 0 {
		query = query.Where("price <= ?", maxPrice)
	}

	// Đếm tổng số dòng
	if err := query.Count(&p.Total).Error; err != nil {
		return nil, err
	}

	var currentOffset int
	if originalTotal > 0 && originalTotal < p.Total {
		currentOffset = p.Offset + int(p.Total-originalTotal)
	} else {
		currentOffset = p.Offset
	}
	// Truy vấn có paging
	result := query.
		Limit(p.Limit).
		Offset(currentOffset).
		Order("created_at DESC").
		Find(&items)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return items, errors.New("ITEM NOT FOUND")
		} else {
			return items, result.Error
		}
	}
	return items, nil
}
