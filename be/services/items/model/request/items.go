package request

import (
	"mime/multipart"
)

type ItemCreationRequest struct {
	Name          string                `form:"name"` // form-data
	Description   string                `form:"description"`
	Quantity      int64                 `form:"quantity"`
	Price         float64               `form:"price"`
	PictureHeader *multipart.FileHeader `form:"picture"`
	PictureFile   *multipart.File
}

type ItemUpdatingRequest struct {
	Name          string                `form:"name"`
	Description   string                `form:"description"`
	Quantity      int64                 `form:"quantity"`
	Price         float64               `form:"price"`
	PictureHeader *multipart.FileHeader `form:"picture"`
	PictureFile   *multipart.File
}

type ItemUploadRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Quantity    string `json:"quantity"`
	Price       string `json:"price"`
	Picture     string `json:"picture"`
	View        string `json:"view"`
	Recommend   string `json:"recommend"`
}
