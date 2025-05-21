package dto

import (
	"time"

	"github.com/google/uuid"
)

type ItemCreationRequest struct {
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description"`
	Quantity    int     `json:"quantity"`
	Price       float64 `json:"price"`
	Picture     string  `json:"picture"`
}

type ItemCreationResponse struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description"`
	Quantity    int       `json:"quantity"`
	Price       float64   `json:"price"`
	Picture     string    `json:"picture"`
	CreatedAt   time.Time `json:"created_at"`
}

type ItemUpdatingRequest struct {
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description"`
	Quantity    int     `json:"quantity"`
	Price       float64 `json:"price"`
	Picture     string  `json:"picture"`
}

type ItemUpdatingResponse struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Quantity    int       `json:"quantity"`
	Price       float64   `json:"price"`
	Picture     string    `json:"picture"`
	UpdatedAt   time.Time `json:"updated_at"`
}
