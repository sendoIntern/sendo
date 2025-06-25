package response

import (
	"time"

	"github.com/google/uuid"
)

type ItemCreationResponse struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Quantity    int64     `json:"quantity"`
	Price       float64   `json:"price"`
	Picture     string    `json:"picture"`
	CreatedAt   time.Time `json:"created_at"`
}

type ItemUpdatingResponse struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Quantity    int64     `json:"quantity"`
	Price       float64   `json:"price"`
	Picture     string    `json:"picture"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type ItemUploadResponse struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Quantity    string `json:"quantity"`
	Price       string `json:"price"`
	Picture     string `json:"picture"`
	View        string `json:"view"`
	Recommend   string `json:"recommend"`
	ItemErr     string `json:"item_err"`
}
