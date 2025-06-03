package usecase

import (
	"be/pkg/utils"
	"be/services/items/model/entity"
	"be/services/items/model/request"
	"be/services/items/repository"

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
