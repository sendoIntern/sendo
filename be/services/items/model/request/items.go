package request

import (
	"mime/multipart"
)

type ItemCreationRequest struct {
	Name        string                `form:"name"` // form-data
	Description string                `form:"description"`
	Quantity    int64                 `form:"quantity"`
	Price       float64               `form:"price"`
	Picture     *multipart.FileHeader `form:"picture"`
}

type ItemUpdatingRequest struct {
	Name        string                `form:"name"`
	Description string                `form:"description"`
	Quantity    int64                 `form:"quantity"`
	Price       float64               `form:"price"`
	Picture     *multipart.FileHeader `form:"picture"`
}
