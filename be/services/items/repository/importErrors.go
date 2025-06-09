package repository

import (
	"be/pkg/db"
	"be/services/items/model/entity"
)

func SaveError(importErr entity.ImportError) error {
	if err := db.DB.Create(&importErr).Error; err != nil {
		return err
	}
	return nil
}

func DeleteError() error {
	result := db.DB.Exec("")
	if result.Error != nil {
		return result.Error
	}
	return nil
}
