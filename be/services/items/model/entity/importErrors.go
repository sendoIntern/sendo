package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ImportError struct {
	ID          uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
	Description string    `json:"description" gorm:"type:text"`
	CreatedAt   time.Time `json:"created_at"`
}

func (ImportErrors *ImportError) BeforeCreate(tx *gorm.DB) (err error) {
	ImportErrors.ID = uuid.New()
	return
}
