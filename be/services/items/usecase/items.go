package usecase

import (
	"be/pkg/db"
	"be/pkg/utils"
	"be/services/items/model/entity"
	"be/services/items/model/request"
	"be/services/items/repository"
	"errors"

	"github.com/google/uuid"
)

func CreateItem(req request.ItemCreationRequest) (entity.Item, error) {
	imageURL, err := utils.UploadToCloudinary(*req.PictureFile, req.PictureHeader)
	if err != nil {
		return entity.Item{}, err
	}
	item := entity.Item{
		Name:        req.Name,
		Description: req.Description,
		Quantity:    req.Quantity,
		Price:       req.Price,
		Picture:     imageURL,
	}

	if err := repository.CreateItem(item); err != nil {
		return entity.Item{}, err
	}

	return item, nil
}

func DeleteItem(id string) error {
	uid, err := uuid.Parse(id)
	if err != nil {
		return err
	}
	return repository.DeleteItem(uid)
}

func UpdateItem(id string, req request.ItemUpdatingRequest) (entity.Item, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return entity.Item{}, err
	}

	var item entity.Item
	item, exist := repository.IsExistItem(uid)
	if !exist {
		return entity.Item{}, errors.New("item not found")
	}
	if req.PictureFile != nil {
		imgURL, err := utils.UploadToCloudinary(*req.PictureFile, req.PictureHeader)
		if err != nil {
			return entity.Item{}, err
		} else {
			item.Picture = imgURL
		}
	}
	item.Name = req.Name
	item.Description = req.Description
	item.Quantity = req.Quantity
	item.Price = req.Price

	err = repository.UpdateItem(item)
	if err != nil {
		return entity.Item{}, err
	}
	return item, nil
}

func GetAllItems() ([]entity.Item, error) {
	var items []entity.Item
	result := db.DB.Find(&items)
	if result.Error != nil {
		return nil, result.Error
	}
	return items, nil
}

func GetItemById(itemId string) (*entity.Item, error) {
	var item entity.Item
	result := db.DB.First(&item, "id = ?", itemId)
	if result.Error != nil {
		return nil, result.Error
	}
	db.DB.Model(&item).Update("view", item.View+1)
	return &item, nil
}
