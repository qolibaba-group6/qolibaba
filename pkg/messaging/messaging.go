package messaging

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/streadway/amqp"
	"log"
)

const (
	rabbitmqURL       = "amqp://guest:guest@localhost:5672/"
	redisCacheTimeout = 3600
	TourQueue         = "tourQueue"
	ClaimQueue        = "claimsQueue"
)

type Messaging struct {
	channel     *amqp.Channel
	redisClient *redis.Client
}

func NewMessaging(channel *amqp.Channel, redisClient *redis.Client) *Messaging {
	return &Messaging{
		channel:     channel,
		redisClient: redisClient,
	}
}

func ConnectToRabbitMQ() (*amqp.Connection, *amqp.Channel, error) {
	conn, err := amqp.Dial(rabbitmqURL)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to connect to RabbitMQ: %w", err)
	}

	channel, err := conn.Channel()
	if err != nil {
		conn.Close()
		return nil, nil, fmt.Errorf("failed to create channel: %w", err)
	}

	return conn, channel, nil
}

// PublishMessage sends a message to a specified RabbitMQ queue
func (m *Messaging) PublishMessage(queueName string, message interface{}) error {
	body, err := json.Marshal(message)
	if err != nil {
		return fmt.Errorf("error marshalling message: %w", err)
	}

	err = m.channel.Publish(
		"",        // exchange
		queueName, // routing key
		false,     // mandatory
		false,     // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
	if err != nil {
		return fmt.Errorf("error publishing message: %w", err)
	}
	log.Printf("Message published to queue: %s", queueName)
	return nil
}

// StartConsumer consumes messages from a specified RabbitMQ queue
func (m *Messaging) StartConsumer(queueName string, handler func(amqp.Delivery)) error {
	msgs, err := m.channel.Consume(
		queueName, // queue
		"",        // consumer
		false,     // auto-ack
		false,     // exclusive
		false,     // no-local
		false,     // no-wait
		nil,       // args
	)
	if err != nil {
		return fmt.Errorf("failed to register a consumer: %w", err)
	}

	for msg := range msgs {
		handler(msg)
	}
	return nil
}

// CacheDataInRedis caches data in Redis with an expiration time
func (m *Messaging) CacheDataInRedis(key string, data interface{}) error {
	ctx := context.Background()
	dataJSON, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("error marshalling data to JSON: %w", err)
	}

	err = m.redisClient.Set(ctx, key, dataJSON, redisCacheTimeout).Err()
	if err != nil {
		return fmt.Errorf("error caching data in Redis: %w", err)
	}

	log.Printf("Data cached in Redis with key: %s", key)
	return nil
}

// HandleClaim processes the claim message and acknowledges it
func (m *Messaging) HandleClaim(msg amqp.Delivery) {
	var claim map[string]interface{}
	err := json.Unmarshal(msg.Body, &claim)
	if err != nil {
		log.Printf("Error unmarshalling claim: %v", err)
		msg.Nack(false, true)
		return
	}

	// Process the claim (e.g., pass to bank service)
	log.Printf("Processed claim: %v", claim)

	// Acknowledge the message to RabbitMQ
	if err := msg.Ack(false); err != nil {
		log.Printf("Error acknowledging message: %v", err)
	}
}

// HandleHotelOffer processes hotel offer messages and caches them in Redis
func (m *Messaging) HandleHotelOffer(msg amqp.Delivery) {
	var hotelOffer map[string]interface{}
	err := json.Unmarshal(msg.Body, &hotelOffer)
	if err != nil {
		log.Printf("Error unmarshalling hotel offer: %v", err)
		msg.Nack(false, true)
		return
	}

	// Cache the hotel offer in Redis
	err = m.CacheDataInRedis(fmt.Sprintf("hotel_offer:%s", hotelOffer["id"]), hotelOffer)
	if err != nil {
		log.Printf("Error caching hotel offer: %v", err)
		msg.Nack(false, true)
		return
	}

	log.Printf("Successfully cached hotel offer: %v", hotelOffer)

	// Acknowledge the message to RabbitMQ
	if err := msg.Ack(false); err != nil {
		log.Printf("Error acknowledging message: %v", err)
	}
}
