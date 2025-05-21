package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Item struct {
	ID          uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
	Name        string    `json:"name" gorm:"unique;not null"`
	Description string    `json:"description" gorm:"type:text"`
	Quantity    int       `json:"quantity" gorm:"default:0"`
	Price       float64   `json:"price" gorm:"default:0"`
	Picture     string    `json:"picture"`
	View        int64     `json:"view" gorm:"default:0"`
	Recommend   int       `json:"recommend" gorm:"default:0"`
	CreateAt    time.Time `json:"create_at"`
	UpdateAt    time.Time `json:"update_at"`
}

// auto generate uuid before save a new user into database
func (item *Item) BeforeCreate(tx *gorm.DB) (err error) {
	item.ID = uuid.New()
	return
}
