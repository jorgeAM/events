package models

import (
	"time"

	"github.com/globalsign/mgo/bson"
)

// Event models
type Event struct {
	ID        bson.ObjectId `json:"id" bson:"_id"`
	Name      string
	Available bool
	StartAt   time.Time
}
