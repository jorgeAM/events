package listener

import (
	"log"

	"github.com/jorgeAM/events/bookingService/db"
	"github.com/jorgeAM/events/bookingService/models"
	"github.com/jorgeAM/events/msgbroker"
	"github.com/jorgeAM/events/msgbroker/events"
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
			ep.handleEvent(event)
		case err = <-errorChan:
			log.Printf("received error while processing msg: %s", err)
		}
	}
}

func (ep *EventProccess) handleEvent(event msgbroker.Event) {
	switch ev := event.(type) {
	case *events.EventCreated:
		ep.DBHandler.SaveEvent(models.Event{ID: ev.ID, Name: ev.Name, Start: ev.Start})
		log.Printf("event %s created %T", ev.ID, ev)
	default:
		log.Println("ÑO")
	}
}
