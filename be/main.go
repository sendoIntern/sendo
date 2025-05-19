package main

import (
	"be/db"
)

func main() {
	sql := db.New()
	sql.Connect()
	defer sql.Close()
}
