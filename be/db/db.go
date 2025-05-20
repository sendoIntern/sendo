package db

import (
	"fmt"
	"log"
	"os"

	"be/entity"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// New một instance mới của DB với giá trị từ .env
func New() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Could not load .env file")
	}

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))

	var err2 error
	DB, err2 = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err2 != nil {
		log.Fatal("Failed to connect to database")
		return
	}

	// Tự động tạo bảng nếu chưa có
	err = DB.AutoMigrate(&entity.User{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
		return
	}

	log.Println("Connect database successful!")
}

// đóng kết nối database
func Close() {
	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatal("Failed to get database instance")
		return
	}
	sqlDB.Close()
}
