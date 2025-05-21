package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Item struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey"`
	Name        string    `gorm:"unique;not null"`
	Description string    `gorm:"type:text"`
	Quantity    int       `gorm:"default:0"`
	Price       float64   `gorm:"default:0"`
	Picture     string
	View        int64 `gorm:"default:0"`
	Recommend   int   `gorm:"default:0"`
	CreateAt    time.Time
	UpdateAt    time.Time
}

// auto generate uuid before save a new user into database
func (item *Item) BeforeCreate(tx *gorm.DB) (err error) {
	item.ID = uuid.New()
	return
}
