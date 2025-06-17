package usecase

import (
	"be/pkg/cloudinary"
	"be/pkg/db"
	"be/pkg/pagination"
	"be/services/items/model/entity"
	"be/services/items/model/request"
	"be/services/items/repository"
	"errors"
	"log"
	"math"
	"mime/multipart"
	"strconv"
	"time"

	"github.com/xuri/excelize/v2"

	"github.com/google/uuid"
)

func CreateItem(req request.ItemCreationRequest) (entity.Item, error) {
	imageURL, err := cloudinary.UploadToCloudinary(*req.PictureFile, req.PictureHeader)
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
		imgURL, err := cloudinary.UploadToCloudinary(*req.PictureFile, req.PictureHeader)
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

func GetAllItems(p *pagination.Paging) ([]entity.Item, error) {
	p.Offset = (p.Page - 1) * p.Limit
	items, err := repository.FetchItems(p)
	if err != nil {
		return items, err
	}
	p.TotalPages = int(math.Ceil(float64(p.Total) / float64(p.Limit)))
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

func ParseExcel(file *multipart.FileHeader) ([]entity.Item, []error) {
	var items []entity.Item
	var errs []error

	f, err := file.Open()
	if err != nil {
		log.Printf("Cannot open file: %v\n", err)
		errs = append(errs, errors.New("Cannot open file excel: "+err.Error()))
		return items, errs
	}
	defer f.Close()

	excelFile, err := excelize.OpenReader(f)
	if err != nil {
		log.Printf("Invalid excel file: %v\n", err)
		errs = append(errs, errors.New("Invalid file type: "+err.Error()))
		return items, errs
	}

	rows, err := excelFile.GetRows("Sheet1")
	if err != nil {
		log.Printf("Cannot read sheet: %v\n", err)
		errs = append(errs, errors.New("Cannot read sheet: "+err.Error()))
		return items, errs
	}

	for i, row := range rows {
		if i == 0 {
			continue // skip header
		}

		// Kiểm tra số lượng cột
		if len(row) < 6 {
			log.Printf("Row %d: Invalid number of columns, expected at least 6, got %d", i, len(row))
			errs = append(errs, errors.New("Invalid columns at row:"+string(rune(i))))
			continue
		}

		price, _ := strconv.ParseFloat(row[3], 64)
		quantity, _ := strconv.ParseInt(row[2], 10, 64)
		view, _ := strconv.ParseInt(row[5], 10, 64)

		// Mặc định recommend = 0 nếu không có cột 7
		recommend := int64(0)
		if len(row) > 6 {
			recommend, _ = strconv.ParseInt(row[6], 10, 64)
		}

		item := entity.Item{
			Name:        row[0],
			Description: row[1],
			Quantity:    quantity,
			Price:       price,
			Picture:     row[4],
			View:        view,
			Recommend:   recommend,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}

		items = append(items, item)
	}

	return items, nil
}

func GetErrorItems() ([]entity.ImportError, error) {
	var errs []entity.ImportError
	result := db.DB.Find(&errs)
	if result.Error != nil {
		return nil, result.Error
	}
	err := repository.DeleteError()
	if err != nil {
		return errs, err
	}
	return errs, nil
}

func GetItemDesc() ([]entity.Item, error) {
	// Lấy 3 item có view cao nhất
	result, err := repository.GetTopViewedItems()
	if err != nil {
		return nil, err
	}
	return result, nil
}
