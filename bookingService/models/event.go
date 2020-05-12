package models

import "time"

// Event model
type Event struct {
	ID    string    `json:"id"`
	Name  string    `json:"name"`
	Start time.Time `json:"start"`
}
