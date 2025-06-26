package config

import (
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

var once sync.Once

func loadEnv() {
	once.Do(func() {
		err := godotenv.Load()
		if err != nil {
			log.Println("No .env file found or failed to load, using system env vars")
		}
	})
}

func GetEnv(key, fallback string) string {
	loadEnv()
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
