package db

import (
	"be/pkg/config"
	"be/services/auth/model/entity"
	itemsEntity "be/services/items/model/entity"
	"fmt"
	"log"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	instance *gorm.DB
	once     sync.Once
)

func GetDB() *gorm.DB {
	once.Do(func() {
		dsn := fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			config.GetEnv("DB_HOST", "localhost"),
			config.GetEnv("DB_PORT", "5432"),
			config.GetEnv("DB_USER", "postgres"),
			config.GetEnv("DB_PASSWORD", ""),
			config.GetEnv("DB_NAME", "postgres"),
		)

		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatalf("failed to connect to DB: %v", err)
		}

		instance = db
		// Tự động tạo bảng nếu chưa có
		err = db.AutoMigrate(
			&entity.User{},
			&itemsEntity.Item{},
			&itemsEntity.ImportError{},
		)
		if err != nil {
			log.Fatal("Failed to migrate database:", err)
			return
		}
	})

	return instance
}
