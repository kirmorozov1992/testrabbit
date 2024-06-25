package testrabbit

import (
	"context"
	"encoding/json"
	"log"

	"git.ipc/kirillmorozov/testrabbit/internal/message"
	amqp "github.com/rabbitmq/amqp091-go"
)

type TestRabbit struct {
	ConnectionStr string
	QName         string
	TxID          int
}

func NewTestRabbit(connectionStr string, qname string, t int) *TestRabbit {
	return &TestRabbit{
		ConnectionStr: connectionStr,
		QName:         qname,
		TxID:          t,
	}
}

func FailOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func (t *TestRabbit) NewConn() *amqp.Connection {
	conn, err := amqp.Dial(t.ConnectionStr)
	FailOnError(err, "Failed to connect to RabbitMQ")
	return conn
}

func (t *TestRabbit) DeleteTransactionFromQueue() {
	msg := message.NewMessage(t.TxID)

	conn := t.NewConn()
	defer conn.Close()

	ch, err := conn.Channel()
	FailOnError(err, "Failed to open a channel")
	defer ch.Close()

	msgs, err := ch.Consume(
		t.QName,
		"",
		false,
		false,
		false,
		false,
		nil,
	)

	FailOnError(err, "Failed to register a consumer")

	unmMsg := message.Message{}

	for m := range msgs {
		err := json.Unmarshal(m.Body, &unmMsg)
		FailOnError(err, "Unmarshal error")

		if unmMsg.TransactionID == msg.TransactionID {
			err := m.Ack(false)
			FailOnError(err, "Error acknowledging message")
			break
		}
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case m, ok := <-msgs:
				if !ok {
					return
				}
				err := ch.Publish(
					"",
					t.QName,
					false,
					false,
					amqp.Publishing{
						ContentType: m.ContentType,
						Body:        m.Body,
					})
				FailOnError(err, "Failed to publish a message")
			}
		}
	}()
}
