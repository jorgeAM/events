package server

import (
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/jorgeAM/events/bookingService/db"
	"github.com/jorgeAM/events/bookingService/routes"
)

// Listen initialize HTTP server
func Listen(endpoint string, dbhandler db.DatabaseHandler) error {
	r := routes.InitBookingRoutes(dbhandler)
	server := handlers.CORS()(r)
	return http.ListenAndServe(endpoint, server)
}
