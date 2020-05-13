package db

import (
	"encoding/json"
	"sync"

	"github.com/jinzhu/gorm"
	// driver to connect with postgres
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/jorgeAM/events/bookingService/models"
)

var (
	sqlLayer *DBLayer
	once     sync.Once
	e        error
)

// DBLayer handle connection
type DBLayer struct {
	DB *gorm.DB
}

// NewSQLLayer handle connection with database
func NewSQLLayer(engine string, url string) (*DBLayer, error) {
	once.Do(func() {
		db, err := gorm.Open(engine, url)
		sqlLayer = &DBLayer{
			DB: db,
		}

		e = err
	})

	return sqlLayer, e
}

// SaveBooking saves booking
func (dbLayer *DBLayer) SaveBooking(booking *models.Booking) ([]byte, error) {
	err := dbLayer.DB.Create(booking).Error

	if err != nil {
		return nil, err
	}

	return json.Marshal(booking)
}

// SaveEvent saves event
func (dbLayer *DBLayer) SaveEvent(event models.Event) error {
	return dbLayer.DB.Create(&event).Error
}

// FindEventByID retrieves an event
func (dbLayer *DBLayer) FindEventByID(id string) (*models.Event, error) {
	var event models.Event
	err := dbLayer.DB.Where("id = ?", id).First(&event).Error

	if err != nil {
		return nil, err
	}

	return &event, nil
}
