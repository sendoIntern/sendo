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

var database = db.GetDB()

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
		Status:      true,
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

func GetAllItems(
	p *pagination.Paging,
	search string,
	minPrice float64,
	maxPrice float64,
	originalTotal int64, 
 ) ([]entity.Item, error) {

	p.Offset = (p.Page - 1) * p.Limit
	items, err := repository.FetchItems(p, search, minPrice, maxPrice, originalTotal)
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
			errs = append(errs, errors.New("ROW "+string(rune(i))+": Invalid number of columns, expected at least 6"))
			continue
		}
		// validate fields
		name := row[0]
		if repository.IsExistItemByName(name) {
			errs = append(errs, errors.New("ROW "+string(rune(i))+": Duplicate name with an existed item"))
			continue
		}
		price, err := strconv.ParseFloat(row[3], 64)
		if err != nil || price < 0 {
			errs = append(errs, errors.New("ROW "+string(rune(i))+": Invalid value at price column, expected positive float number"))
			continue
		}
		quantity, err := strconv.ParseInt(row[2], 10, 64)
		if err != nil || quantity < 0 {
			errs = append(errs, errors.New("ROW "+string(rune(i))+": Invalid value at quantity column, expected positive integer number"))
			continue
		}
		view, err := strconv.ParseInt(row[5], 10, 64)
		if err != nil || view < 0 {
			errs = append(errs, errors.New("ROW "+string(rune(i))+": Invalid value at view column, expected positive integer number"))
			continue
		}

		// Mặc định recommend = 0 nếu không có cột 7
		recommend := int64(0)
		if len(row) > 6 {
			recommend, err = strconv.ParseInt(row[6], 10, 64)
			if err != nil || recommend < 0 {
				errs = append(errs, errors.New("ROW "+string(rune(i))+": Invalid value at recommend column, expected positive integer number"))
				continue
			}
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

	return items, errs
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
