package dto

import (
	"mime/multipart"
	"time"

	"github.com/google/uuid"
)

type ItemCreationRequest struct {
	Name        string                `form:"name"` // form-data
	Description string                `form:"description"`
	Quantity    int64                 `form:"quantity"`
	Price       float64               `form:"price"`
	Picture     *multipart.FileHeader `form:"picture"`
}

type ItemCreationResponse struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Quantity    int64     `json:"quantity"`
	Price       float64   `json:"price"`
	Picture     string    `json:"picture"`
	CreatedAt   time.Time `json:"created_at"`
}

type ItemUpdatingRequest struct {
	Name        string                `form:"name"`
	Description string                `form:"description"`
	Quantity    int64                 `form:"quantity"`
	Price       float64               `form:"price"`
	Picture     *multipart.FileHeader `form:"picture"`
}

type ItemUpdatingResponse struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Quantity    int64     `json:"quantity"`
	Price       float64   `json:"price"`
	Picture     string    `json:"picture"`
	UpdatedAt   time.Time `json:"updated_at"`
}
