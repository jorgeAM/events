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
	sqlLayer *dbLayer
	once     sync.Once
	e        error
)

type dbLayer struct {
	DB *gorm.DB
}

// NewSQLLayer handle connection with database
func NewSQLLayer(engine string, url string) (*dbLayer, error) {
	once.Do(func() {
		db, err := gorm.Open(engine, url)
		sqlLayer = &dbLayer{
			DB: db,
		}

		e = err
	})

	return sqlLayer, e
}

func (dbLayer *dbLayer) SaveBooking(booking *models.Booking) ([]byte, error) {
	err := dbLayer.DB.Create(booking).Error

	if err != nil {
		return nil, err
	}

	return json.Marshal(booking)
}

func (dbLayer *dbLayer) SaveEvent(event models.Event) error {
	return dbLayer.DB.Create(&event).Error
}
