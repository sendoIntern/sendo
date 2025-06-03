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
