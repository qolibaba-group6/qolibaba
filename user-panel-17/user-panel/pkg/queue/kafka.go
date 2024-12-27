package queue

import (
	"context"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

var writer *kafka.Writer
var reader *kafka.Reader

// InitializeKafka initializes Kafka producer and consumer
func InitializeKafka(broker string, topic string) {
	// Initialize Kafka writer (Producer)
	writer = &kafka.Writer{
		Addr:     kafka.TCP(broker),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}

	// Initialize Kafka reader (Consumer)
	reader = kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{broker},
		Topic:   topic,
		GroupID: "group-id",
	})
}

// ProduceMessage sends a message to Kafka
func ProduceMessage(key, value string) error {
	err := writer.WriteMessages(context.Background(),
		kafka.Message{
			Key:   []byte(key),
			Value: []byte(value),
		},
	)
	if err != nil {
		log.Printf("Failed to write message to Kafka: %v", err)
		return err
	}
	log.Printf("Message sent to Kafka: key=%s value=%s", key, value)
	return nil
}

// ConsumeMessages reads messages from Kafka
func ConsumeMessages(handler func(key, value string)) {
	for {
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Printf("Failed to read message from Kafka: %v", err)
			break
		}
		log.Printf("Message received from Kafka: key=%s value=%s", string(msg.Key), string(msg.Value))
		handler(string(msg.Key), string(msg.Value))
	}
}

// CloseKafka closes Kafka producer and consumer
func CloseKafka() {
	writer.Close()
	reader.Close()
}
