package main

import (
	"log"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/jorgeAM/events/bookingService/db"
	"github.com/jorgeAM/events/bookingService/server"
	"github.com/jorgeAM/events/msgbroker/rabbitmq"
	"github.com/streadway/amqp"
)

func main() {
	endpoint := os.Getenv("ENDPOINT")
	database := os.Getenv("DATABASE_ENGINE")
	databaseURL := os.Getenv("DATABASE_URL")
	amqpURL := os.Getenv("AMQP_URL")

	log.Println("connecting to database")
	dbhandler, err := db.NewSQLLayer(database, databaseURL)
	defer dbhandler.DB.Close()

	if err != nil {
		log.Fatal(err)
	}

	log.Println("connecting to message broker")
	conn, err := amqp.Dial(amqpURL)

	if err != nil {
		log.Fatal(err)
	}

	_, err = rabbitmq.NewAmqpEventListener(conn, "events")

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Server starts running")
	log.Fatal(server.Listen(endpoint))
}
