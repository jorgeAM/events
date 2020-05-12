package server

import (
	"net/http"

	"github.com/jorgeAM/events/eventService/db"
	"github.com/jorgeAM/events/eventService/routes"
	"github.com/jorgeAM/events/msgbroker"
)

// Listen initialize HTTP server
func Listen(endpoint string, dbHandler db.DatabaseHandler, emitter msgbroker.EventEmitter) error {
	r := routes.InitEventRoutes(dbHandler, emitter)
	return http.ListenAndServe(endpoint, r)
}
