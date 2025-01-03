// pkg/messaging/rabbitmq.go
package messaging

import (
	"log"

	"github.com/streadway/amqp"
)

// RabbitMQ implements the Messaging interface using RabbitMQ
type RabbitMQ struct {
	Conn    *amqp.Connection
	Channel *amqp.Channel
	Queue   amqp.Queue
}

// NewRabbitMQ creates a new RabbitMQ instance
func NewRabbitMQ(uri, queueName string) *RabbitMQ {
	conn, err := amqp.Dial(uri)
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
	}

	q, err := ch.QueueDeclare(
		queueName, // name
		true,      // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		log.Fatalf("Failed to declare a queue: %v", err)
	}

	return &RabbitMQ{
		Conn:    conn,
		Channel: ch,
		Queue:   q,
	}
}

// Publish sends a message to the specified queue
func (r *RabbitMQ) Publish(queueName string, message string) error {
	return r.Channel.Publish(
		"",        // exchange
		queueName, // routing key
		false,     // mandatory
		false,     // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		},
	)
}

// Consume listens to the specified queue and handles incoming messages
func (r *RabbitMQ) Consume(queueName string, handler func(message string) error) error {
	msgs, err := r.Channel.Consume(
		queueName, // queue
		"",        // consumer
		true,      // auto-ack
		false,     // exclusive
		false,     // no-local
		false,     // no-wait
		nil,       // args
	)
	if err != nil {
		return err
	}

	go func() {
		for d := range msgs {
			if err := handler(string(d.Body)); err != nil {
				log.Printf("Error handling message: %v", err)
			}
		}
	}()

	return nil
}

// Close closes the RabbitMQ connection and channel
func (r *RabbitMQ) Close() error {
	if err := r.Channel.Close(); err != nil {
		return err
	}
	if err := r.Conn.Close(); err != nil {
		return err
	}
	return nil
}
