package models

import (
	"time"
)

// Event model
type Event struct {
	ID    string    `json:"id" gorm:"primary_key"`
	Name  string    `json:"name"`
	Start time.Time `json:"start"`
}
