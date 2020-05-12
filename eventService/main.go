package main

import (
	"fmt"
	"log"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/jorgeAM/events/eventService/db"
	"github.com/jorgeAM/events/eventService/server"
)

func main() {
	endpoint := os.Getenv("ENDPOINT")
	database := os.Getenv("DATABASE_ENGINE")
	databaseURL := os.Getenv("DATABASE_URL")
	fmt.Println(endpoint)
	log.Println("connecting to ", database)
	dbHandler, err := db.NewPersistenceLayer(db.TYPE(database), databaseURL)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Server starts running")
	log.Fatal(server.Listen(endpoint, dbHandler))
}
