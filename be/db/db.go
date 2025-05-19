package db

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
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

func (s *Sql) Connect() {

	dataSource := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		s.Host, s.Port, s.UserName, s.Password, s.DbName)
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
