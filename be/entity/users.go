package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID       uuid.UUID `gorm:"type:uuid;primaryKey"`
	GoogleID string    `gorm:"unique"`
	Name     string
	Email    string `gorm:"unique"`
	Picture  string
}

// auto generate uuid before save a new user into database
func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	user.ID = uuid.New()
	return
}
