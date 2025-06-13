package repository

import (
	"be/pkg/db"
	"be/services/items/model/entity"
	"errors"

	"github.com/google/uuid"
)

func CreateItem(item entity.Item) error {
	if err := db.DB.Create(&item).Error; err != nil {
		return err
	}
	return nil
}

func DeleteItem(id uuid.UUID) error {
	result := db.DB.Delete(&entity.Item{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("item not found")
	}
	return nil
}

func UpdateItem(item entity.Item) error {
	if err := db.DB.Save(&item).Error; err != nil {
		return err
	}
	return nil
}

func IsExistItem(id uuid.UUID) (entity.Item, bool) {
	var item entity.Item
	result := db.DB.First(&item, "id = ?", id)
	if result.Error != nil {
		return entity.Item{}, false
	}
	return item, true
}

// lấy 3 item có view cao nhất
func GetTopViewedItems() ([]entity.Item, error) {
	var items []entity.Item
	if err := db.DB.Order("view DESC").Limit(3).Find(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}