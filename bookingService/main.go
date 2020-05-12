package main

import (
	"log"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/jorgeAM/events/bookingService/server"
)

func main() {
	endpoint := os.Getenv("ENDPOINT")

	log.Println("Server starts running")
	log.Fatal(server.Listen(endpoint))
}
