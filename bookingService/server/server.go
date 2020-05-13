package server

import (
	"net/http"

	"github.com/jorgeAM/events/bookingService/db"
	"github.com/jorgeAM/events/bookingService/routes"
)

// Listen initialize HTTP server
func Listen(endpoint string, dbhandler db.DatabaseHandler) error {
	r := routes.InitBookingRoutes(dbhandler)
	return http.ListenAndServe(endpoint, r)
}
