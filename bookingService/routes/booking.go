package routes

import (
	"github.com/gorilla/mux"
	"github.com/jorgeAM/events/bookingService/db"
	"github.com/jorgeAM/events/bookingService/handler"
)

// InitBookingRoutes init booking routes
func InitBookingRoutes(dbhandler db.DatabaseHandler) *mux.Router {
	r := mux.NewRouter()
	bookingHandler := handler.NewBookingHandler(dbhandler)
	r.HandleFunc("/events/{id}/booking", bookingHandler.SaveBookingHandler).Methods("POST")

	return r
}
