package msgbroker

// EventListener listen events
type EventListener interface {
	Listen(name string) (<-chan Event, <-chan error, error)
}
