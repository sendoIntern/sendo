package db

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type Sql struct {
	Db       *sqlx.DB
	Host     string
	Port     string
	UserName string
	Password string
	DbName   string
}

// New một instance mới của Sql với giá trị từ .env
func New() *Sql {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Could not load .env file")
	}

	return &Sql{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		UserName: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DbName:   os.Getenv("DB_NAME"),
	}
}

func (s *Sql) Connect() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	dataSource := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))
	s.Db = sqlx.MustConnect("postgres", dataSource)

	if err := s.Db.Ping(); err != nil {
		log.Fatal("Cannot connect to database")
		return
	}

	log.Println("Connect database successful!")
}

func (s *Sql) Close() {
	s.Db.Close()
}
