package events

// EventCreated represents message to emit
type EventCreated struct {
	ID   string
	Name string
}

// EventName return name of event
func (ev *EventCreated) EventName() string {
	return "event.created"
}
