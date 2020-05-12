package main

import (
	"log"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/jorgeAM/events/eventService/db"
	"github.com/jorgeAM/events/eventService/server"
	"github.com/jorgeAM/events/msgbroker/rabbitmq"
	"github.com/streadway/amqp"
)

func main() {
	endpoint := os.Getenv("ENDPOINT")
	database := os.Getenv("DATABASE_ENGINE")
	databaseURL := os.Getenv("DATABASE_URL")
	amqpURL := os.Getenv("AMQP_URL")

	log.Println("connecting to ", database)
	dbHandler, err := db.NewPersistenceLayer(db.TYPE(database), databaseURL)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("connecting to message broker")
	conn, err := amqp.Dial(amqpURL)

	if err != nil {
		log.Fatal(err)
	}

	emitter, err := rabbitmq.NewAmqpEventEmitter(conn)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Server starts running")
	log.Fatal(server.Listen(endpoint, dbHandler, emitter))
}
