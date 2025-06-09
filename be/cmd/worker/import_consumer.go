package main

import (
	"encoding/json"
	"log"
	"os"

	"be/pkg/db"
	"be/services/items/model/entity"
	"be/services/items/repository"

	"github.com/joho/godotenv"
	"github.com/streadway/amqp"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Could not load .env file", err)
	}

	db.New()
	defer db.Close()

	conn, _ := amqp.Dial(os.Getenv("RABBITMQ_URL"))
	ch, _ := conn.Channel()
	q, _ := ch.QueueDeclare(
		"item_import",
		false,
		false,
		false,
		false,
		nil,
	)

	msgs, _ := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	log.Print("Consumer is running...")

	var importErr entity.ImportError
	for msg := range msgs {
		var item entity.Item
		if err := json.Unmarshal(msg.Body, &item); err != nil {
			log.Printf("Error parse message: %v", err)
			importErr.Description = "Cannot unmarshal from message: " + err.Error()
			repository.SaveError(importErr)
			continue
		}

		if err := repository.CreateItem(item); err != nil {
			log.Printf("Cannot create item: %v", err)
			importErr.Description = "Cannot create item: " + err.Error()
			repository.SaveError(importErr)
			continue
		}
		log.Printf("Create item success: %s", item.Name)
	}

}
