package usecase

import (
	"be/pkg/cloudinary"
	"be/pkg/db"
	"be/pkg/pagination"
	"be/services/items/model/entity"
	"be/services/items/model/request"
	"be/services/items/model/response"
	"be/services/items/repository"
	"errors"
	"log"
	"math"
	"mime/multipart"
	"strconv"

	"github.com/xuri/excelize/v2"

	"github.com/google/uuid"
)

var database = db.GetDB()

func CreateItem(req request.ItemCreationRequest) (entity.Item, error) {
	item := entity.Item{
		Name:        req.Name,
		Description: req.Description,
		Quantity:    req.Quantity,
		Price:       req.Price,
		Status:      true,
	}
	var imageURL string
	var err error
	if req.PictureHeader == nil || req.PictureFile == nil {
		imageURL = "https://res.cloudinary.com/dngsic1a0/image/upload/v1750912253/no-image-dafault.png.png"
	} else {
		imageURL, err = cloudinary.UploadToCloudinary(*req.PictureFile, req.PictureHeader)
		if err != nil {
			return entity.Item{}, err
		}
	}
	item.Picture = imageURL

	fieldErr := ValidateItem(item)
	if fieldErr != nil {
		return entity.Item{}, fieldErr
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

func ChangeStatusItem(id string) (entity.Item, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return entity.Item{}, err
	}
	return repository.ChangeStatusItem(uid)
}

func UpdateItem(id string, req request.ItemUpdatingRequest) (entity.Item, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return entity.Item{}, err
	}

	var item entity.Item
	item, exist := repository.IsExistItem(uid)
	if !exist {
		return entity.Item{}, errors.New("ITEM NOT FOUND")
	}
	if req.PictureFile != nil {
		imgURL, err := cloudinary.UploadToCloudinary(*req.PictureFile, req.PictureHeader)
		if err != nil {
			return entity.Item{}, err
		} else {
			item.Picture = imgURL
		}
	}
	if repository.IsExistItemByName(req.Name) && item.Name != req.Name {
		return entity.Item{}, errors.New("THIS NAME IS EXISTED")
	}
	item.Name = req.Name
	item.Description = req.Description
	item.Quantity = req.Quantity
	item.Price = req.Price
	if item.Quantity < 0 || item.Price < 0 {
		return entity.Item{}, errors.New("QUANTITY AND PRICE CANNOT NEGATIVE")
	}

	err = repository.UpdateItem(item)
	if err != nil {
		return entity.Item{}, err
	}
	return item, nil
}

func GetAllItems(
	p *pagination.Paging,
	search string,
	minPrice float64,
	maxPrice float64,
	originalTotal int64,
	status string) ([]entity.Item, error) {

	p.Offset = (p.Page - 1) * p.Limit
	items, err := repository.FetchItems(p, search, minPrice, maxPrice, originalTotal, status)
	if err != nil {
		return items, err
	}

	if originalTotal > 0 && originalTotal < p.Total {
		p.Total = originalTotal // original in cache
	}
	p.TotalPages = int(math.Ceil(float64(p.Total) / float64(p.Limit)))
	return items, nil
}

func GetItemById(itemId string) (*entity.Item, error) {
	var item entity.Item
	result := database.First(&item, "id = ?", itemId)
	if result.Error != nil {
		return nil, result.Error
	}
	database.Model(&item).Update("view", item.View+1)
	return &item, nil
}

func ParseExcel(file *multipart.FileHeader) ([]response.ItemUploadResponse, error) {
	var items []response.ItemUploadResponse

	f, err := file.Open()
	if err != nil {
		log.Printf("Cannot open file: %v\n", err)
		return items, errors.New("CANNOT OPEN FILE")
	}
	defer f.Close()

	excelFile, err := excelize.OpenReader(f)
	if err != nil {
		log.Printf("Invalid excel file: %v\n", err)
		return items, errors.New("INVALID EXCEL FILE")
	}

	rows, err := excelFile.GetRows("Sheet1")
	if err != nil {
		log.Printf("Cannot read sheet: %v\n", err)
		return items, errors.New("CANNOT READ SHEET")
	}

	for i, row := range rows {
		if i == 0 {
			continue // skip header
		}

		var item response.ItemUploadResponse
		item.ItemErr = ""
		// Kiểm tra số lượng cột
		if len(row) < 6 {
			item.ItemErr += "Invalid number of columns, expected at least 6___"
		}
		// validate fields
		item.Name = row[0]
		if repository.IsExistItemByName(item.Name) {
			item.ItemErr += "Duplicate name with an existed item___"
		}

		item.Price = row[3]
		price, err := strconv.ParseFloat(item.Price, 64)
		if err != nil || price < 0 {
			item.ItemErr += "Invalid value at price column, expected positive float number___"
		}

		item.Quantity = row[2]
		quantity, err := strconv.ParseInt(item.Quantity, 10, 64)
		if err != nil || quantity < 0 {
			item.ItemErr += "Invalid value at quantity column, expected positive integer number___"
		}

		item.View = row[5]
		view, err := strconv.ParseInt(item.View, 10, 64)
		if err != nil || view < 0 {
			item.ItemErr += "Invalid value at view column, expected positive integer number___"
		}

		// Mặc định recommend = 0 nếu không có cột 7
		item.Recommend = "0"
		if len(row) > 6 {
			item.Recommend = row[6]
			recommend, err := strconv.ParseInt(item.Recommend, 10, 64)
			if err != nil || recommend < 0 {
				item.ItemErr += "Invalid value at recommend column, expected positive integer number___"
			}
		}
		item.Description = row[1]
		item.Picture = row[4]
		items = append(items, item)
	}

	return items, nil
}

func GetErrorItems() ([]entity.ImportError, error) {
	var errs []entity.ImportError
	result := database.Find(&errs)
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

func ValidateItem(item entity.Item) error {
	if repository.IsExistItemByName(item.Name) {
		return errors.New("THIS ITEM NAME IS EXISTED")
	}
	if item.Quantity < 0 || item.Price < 0 {
		return errors.New("QUANTITY AND PRICE CANNOT NEGATIVE")
	}

	return nil
}
