// pkg/messaging/messaging.go
package messaging

// Messaging defines the interface for messaging operations
type Messaging interface {
	// Publish sends a message to a specified queue
	Publish(queueName string, message string) error

	// Consume listens to a specified queue and handles incoming messages
	Consume(queueName string, handler func(message string) error) error
}
