package msgbroker

// Event interface should be implement for all event
type Event interface {
	EventName() string
}
