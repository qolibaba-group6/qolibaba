
package queue

import (
	"log"

	"github.com/segmentio/kafka-go"
)

type KafkaProducer struct {
	writer *kafka.Writer
}

func NewKafkaProducer(brokerAddress string, topic string) *KafkaProducer {
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{brokerAddress},
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	})

	return &KafkaProducer{writer: writer}
}

func (k *KafkaProducer) PublishMessage(key, value string) error {
	err := k.writer.WriteMessages(nil, kafka.Message{
		Key:   []byte(key),
		Value: []byte(value),
	})
	if err != nil {
		log.Printf("Failed to publish message: %v", err)
		return err
	}
	log.Printf("Message published: key=%s, value=%s", key, value)
	return nil
}
