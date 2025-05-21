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
	ID          uuid.UUID
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description"`
	Quantity    int       `json:"quantity"`
	Price       float64   `json:"price"`
	Picture     string    `json:"picture"`
	CreateAt    time.Time `json:"create_at"`
}
