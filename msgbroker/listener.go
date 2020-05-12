package msgbroker

// EventListener listen events
type EventListener interface {
	Listen(eventNames ...string) (<-chan Event, <-chan error, error)
}
