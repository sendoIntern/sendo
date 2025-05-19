package main

import (
	"be/db"

	"log"

	"os"

	"github.com/joho/godotenv"

	"github.com/gofiber/fiber/v2"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	sql := &db.Sql{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		UserName: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DbName:   os.Getenv("DB_NAME"),
	}
	sql.Connect()
	defer sql.Close()
	port := os.Getenv("PORT")
	app := fiber.New()
	app.Listen(":" + port)
}
