package events

import "time"

// EventCreated represents message to emit
type EventCreated struct {
	ID    string    `json:"id"`
	Name  string    `json:"name"`
	Start time.Time `json:"start"`
}

// EventName return name of event
func (ev *EventCreated) EventName() string {
	return "event.created"
}
