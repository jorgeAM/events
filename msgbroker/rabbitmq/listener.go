package rabbitmq

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/jorgeAM/events/msgbroker"
	"github.com/jorgeAM/events/msgbroker/events"
	"github.com/streadway/amqp"
)

type amqpEventListener struct {
	conn  *amqp.Connection
	queue string
}

// NewAmqpEventListener creates a new instance of amqpEventListener
func NewAmqpEventListener(conn *amqp.Connection, queue string) (msgbroker.EventListener, error) {
	listener := &amqpEventListener{
		conn:  conn,
		queue: queue,
	}

	ch, err := listener.conn.Channel()
	defer ch.Close()

	if err != nil {
		return nil, err
	}

	_, err = ch.QueueDeclare(listener.queue, true, false, false, false, nil)

	if err != nil {
		return nil, err
	}

	return listener, nil
}

func (e *amqpEventListener) Listen(eventNames ...string) (<-chan msgbroker.Event, <-chan error, error) {
	ch, err := e.conn.Channel()

	if err != nil {
		return nil, nil, err
	}

	for _, eventName := range eventNames {
		if err := ch.QueueBind(e.queue, eventName, "events", false, nil); err != nil {
			return nil, nil, err
		}
	}

	msgs, err := ch.Consume(e.queue, "", false, false, false, false, nil)

	if err != nil {
		return nil, nil, err
	}

	eventChan := make(chan msgbroker.Event)
	errorChan := make(chan error)

	go func() {
		for msg := range msgs {
			rawEventName, ok := msg.Headers["x-event-name"]

			if !ok {
				errorChan <- errors.New("Event does not have rigth header")
				msg.Nack(false, false)
				continue
			}

			eventName, ok := rawEventName.(string)

			if !ok {
				errorChan <- errors.New("Event's header isn't string")
				msg.Nack(false, false)
				continue
			}

			var event msgbroker.Event

			switch eventName {
			case "event.created":
				event = new(events.EventCreated)
			default:
				errorChan <- fmt.Errorf("event type %s is unknow", eventName)
				continue
			}

			err := json.Unmarshal(msg.Body, event)

			if err != nil {
				errorChan <- err
				continue
			}

			eventChan <- event
			msg.Ack(false)
		}
	}()

	return eventChan, errorChan, nil
}
