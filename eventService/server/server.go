package server

import (
	"net/http"

	"github.com/jorgeAM/events/eventService/db"
	"github.com/jorgeAM/events/eventService/routes"
)

// Listen initialize HTTP server
func Listen(endpoint string, dbHandler db.DatabaseHandler) error {
	r := routes.InitEventRoutes(dbHandler)
	return http.ListenAndServe(endpoint, r)
}
