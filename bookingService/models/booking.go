package models

import "time"

// Booking model
type Booking struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	Start     time.Time `json:"start"`
	Seats     uint      `json:"seats"`
	EventID   string    `json:"eventId" gorm:"column:eventId"`
	CreatedAt time.Time `json:"createAt" gorm:"column:createdAt"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"column:updatedAt"`
}
