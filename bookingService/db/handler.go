package db

import "github.com/jorgeAM/events/bookingService/models"

// DatabaseHandler handles method in database
type DatabaseHandler interface {
	SaveBooking(booking *models.Booking) ([]byte, error)
	SaveEvent(event models.Event) error
}
