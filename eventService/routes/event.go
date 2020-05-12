package routes

import (
	"github.com/gorilla/mux"
	"github.com/jorgeAM/events/eventService/db"
	"github.com/jorgeAM/events/eventService/handlers"
	"github.com/jorgeAM/events/msgbroker"
)

// InitEventRoutes init event routes
func InitEventRoutes(dbHandler db.DatabaseHandler, emitter msgbroker.EventEmitter) *mux.Router {
	r := mux.NewRouter()
	handler := handlers.NewEventHandler(dbHandler, emitter)
	s := r.PathPrefix("/events").Subrouter()
	s.HandleFunc("/", handler.CreateEventHandler).Methods("POST")
	s.HandleFunc("/", handler.ListEventHandler).Methods("GET")
	s.HandleFunc("/{id}", handler.GetEventHandler).Methods("GET")
	return s
}
