package server

import (
	"net/http"

	"github.com/jorgeAM/events/bookingService/routes"
)

// Listen initialize HTTP server
func Listen(endpoint string) error {
	r := routes.InitBookingRoutes()
	return http.ListenAndServe(endpoint, r)
}
