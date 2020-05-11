package db

import (
	"github.com/jorgeAM/events/eventService/db/mongo"
)

// TYPE is a custom type to handle database
type TYPE string

// constants
const (
	MONGODB = "mongodb"
)

// NewPersistenceLayer initialize persistenceLayer
func NewPersistenceLayer(database TYPE, url string) (DatabaseHandler, error) {
	switch database {
	case MONGODB:
		return mongo.NewMongoDBLayer(url)
	default:
		return nil, nil
	}
}
