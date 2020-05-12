package msgbroker

// EventEmitter emits events
type EventEmitter interface {
	Emit(event ...Event) error
}
