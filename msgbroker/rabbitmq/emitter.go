package rabbitmq

import (
	"encoding/json"

	"github.com/jorgeAM/events/msgbroker"
	"github.com/streadway/amqp"
)

type amqpEventEmitter struct {
	conn *amqp.Connection
}

// NewAmqpEventEmitter creates a new instance of amqpEventEmitter
func NewAmqpEventEmitter(conn *amqp.Connection) (msgbroker.EventEmitter, error) {
	emitter := &amqpEventEmitter{
		conn: conn,
	}

	ch, err := emitter.conn.Channel()
	defer ch.Close()

	if err != nil {
		return nil, err
	}

	err = ch.ExchangeDeclare("events", "topic", true, false, false, false, nil)

	if err != nil {
		return nil, err
	}

	return emitter, nil
}

func (e *amqpEventEmitter) Emit(event msgbroker.Event) error {
	bytes, err := json.Marshal(event)

	if err != nil {
		return err
	}

	ch, err := e.conn.Channel()
	defer ch.Close()

	if err != nil {
		return err
	}

	msg := amqp.Publishing{
		Headers:     amqp.Table{"x-event-name": event.EventName()},
		ContentType: "application/json",
		Body:        bytes,
	}

	return ch.Publish("events", event.EventName(), false, false, msg)
}
