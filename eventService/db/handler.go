package db

import "github.com/jorgeAM/events/eventService/models"

// DatabaseHandler handles method in database
type DatabaseHandler interface {
	CreateEvent(event *models.Event) ([]byte, error)
	ListAvailableEvents() ([]models.Event, error)
	GetEvent(id string) (*models.Event, error)
}
