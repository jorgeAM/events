package models

import (
	"time"

	"github.com/globalsign/mgo/bson"
)

// Event models
type Event struct {
	ID        bson.ObjectId `json:"id" bson:"_id"`
	Name      string        `json:"name" bson:"name"`
	Available bool          `json:"available" bson:"available"`
	StartAt   time.Time     `json:"startAt" bson:"startAt"`
}
