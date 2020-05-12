package models

import "time"

// Booking model
type Booking struct {
	ID      uint      `json:"id"`
	Start   time.Time `json:"start"`
	Seats   uint      `json:"seats"`
	EventID string    `json:"eventId"`
}
