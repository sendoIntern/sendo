package repository

import (
	"be/services/items/model/entity"
)

func SaveError(importErr entity.ImportError) error {
	if err := database.Create(&importErr).Error; err != nil {
		return err
	}
	return nil
}

func DeleteError() error {
	result := database.Exec("DELETE from import_errors")
	if result.Error != nil {
		return result.Error
	}
	return nil
}
