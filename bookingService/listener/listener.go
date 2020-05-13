package listener

import (
	"log"

	"github.com/jorgeAM/events/bookingService/db"
	"github.com/jorgeAM/events/msgbroker"
)

// EventProccess handle listener and database
type EventProccess struct {
	Listener  msgbroker.EventListener
	DBHandler db.DatabaseHandler
}

// Proccess proccess events from message broker
func (ep *EventProccess) Proccess() error {
	log.Println("Listening to events...")
	eventChan, errorChan, err := ep.Listener.Listen("event.created")

	if err != nil {
		return err
	}

	for {
		select {
		case event := <-eventChan:
			log.Println("A")
			log.Println(event)
		case err = <-errorChan:
			log.Println("B")
			log.Printf("received error while processing msg: %s", err)
		}
	}
}
