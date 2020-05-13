package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jorgeAM/events/bookingService/db"
	"github.com/jorgeAM/events/bookingService/models"
)

type bookingHandlerService struct {
	dbHandler db.DatabaseHandler
}

// NewBookingHandler creates new instance of bookingHandlerService
func NewBookingHandler(dbhandler db.DatabaseHandler) *bookingHandlerService {
	return &bookingHandlerService{
		dbHandler: dbhandler,
	}
}

func (b *bookingHandlerService) SaveBookingHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	event, err := b.dbHandler.FindEventByID(params["id"])

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Evento no encontrado"))
		return
	}

	var booking models.Booking

	err = json.NewDecoder(r.Body).Decode(&booking)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error al parsear body con booking"))
		return
	}

	booking.EventID = event.ID
	booking.Start = event.Start

	bytes, err := b.dbHandler.SaveBooking(&booking)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error guardar booking"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(bytes)
}
