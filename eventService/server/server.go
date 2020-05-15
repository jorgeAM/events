package server

import (
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/jorgeAM/events/eventService/db"
	"github.com/jorgeAM/events/eventService/routes"
	"github.com/jorgeAM/events/msgbroker"
)

// Listen initialize HTTP server
func Listen(endpoint string, dbHandler db.DatabaseHandler, emitter msgbroker.EventEmitter) error {
	r := routes.InitEventRoutes(dbHandler, emitter)
	server := handlers.CORS()(r)
	return http.ListenAndServe(endpoint, server)
}
