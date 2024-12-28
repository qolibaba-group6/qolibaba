
package queue

import (
	"log"
	"github.com/Shopify/sarama"
)

// KafkaProducer represents a Kafka producer
type KafkaProducer struct {
	producer sarama.SyncProducer
}

// NewKafkaProducer creates a new Kafka producer with Sarama
func NewKafkaProducer(brokers []string) (*KafkaProducer, error) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5
	config.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		return nil, err
	}
	return &KafkaProducer{producer: producer}, nil
}

// SendMessage sends a message to a Kafka topic
func (kp *KafkaProducer) SendMessage(topic, message string) error {
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(message),
	}
	_, _, err := kp.producer.SendMessage(msg)
	if err != nil {
		log.Printf("Failed to send message to Kafka: %v", err)
		return err
	}
	log.Printf("Message sent to Kafka topic %s: %s", topic, message)
	return nil
}
