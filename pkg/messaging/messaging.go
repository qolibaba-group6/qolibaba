package messaging

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/streadway/amqp"
	"log"
	bankPort "qolibaba/internal/bank/port"
	travelPort "qolibaba/internal/travel_agencies/port"
	"qolibaba/pkg/adapter/storage/types"
)

const rabbitmqURL = "amqp://guest:guest@localhost:5672/"
const claimsQueue = "claimsQueue"
const hotelOfferQueue = "hotel_offer_queue"
const redisCacheTimeout = 3600

type Messaging struct {
	channel     *amqp.Channel
	bankService bankPort.Service
	redisClient *redis.Client
	tourService travelPort.Service
}

func NewMessaging(channel *amqp.Channel, bankService bankPort.Service, redisClient *redis.Client, tourService travelPort.Service) *Messaging {
	return &Messaging{
		channel:     channel,
		bankService: bankService,
		redisClient: redisClient,
		tourService: tourService,
	}
}

func ConnectToRabbitMQ() (*amqp.Connection, *amqp.Channel, error) {
	conn, err := amqp.Dial(rabbitmqURL)
	if err != nil {
		return nil, nil, err
	}

	channel, err := conn.Channel()
	if err != nil {
		return nil, nil, err
	}

	return conn, channel, nil
}

func (m *Messaging) CloseConnection() {
	if err := m.channel.Close(); err != nil {
		log.Printf("Error closing channel: %v", err)
	}
}

func (m *Messaging) PublishClaimToBank(claimData []byte) (*uint, error) {
	err := m.channel.Publish(
		"claims_exchange",
		"claims_routing_key",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        claimData,
		},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to publish claim: %w", err)
	}

	// Consume response
	msg, err := m.channel.Consume(
		"claims_response_queue",
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to consume response: %w", err)
	}

	for response := range msg {
		var claimResp struct {
			ClaimID uint `json:"claim_id"`
		}
		err := json.Unmarshal(response.Body, &claimResp)
		if err != nil {
			return nil, fmt.Errorf("error unmarshalling claim response: %w", err)
		}
		return &claimResp.ClaimID, nil
	}

	return nil, fmt.Errorf("no response from Bank Service")
}

func (m *Messaging) StartClaimConsumer() {
	conn, channel, err := ConnectToRabbitMQ()
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer func(conn *amqp.Connection) {
		err := conn.Close()
		if err != nil {
			log.Printf("Error closing connection: %v", err)
		}
	}(conn)
	defer func(channel *amqp.Channel) {
		err := channel.Close()
		if err != nil {
			log.Printf("Error closing channel: %v", err)
		}
	}(channel)

	queue, err := channel.QueueDeclare(
		claimsQueue,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to declare queue: %v", err)
	}

	msgs, err := channel.Consume(
		queue.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to register a consumer: %v", err)
	}

	log.Println("Waiting for claims...")

	for msg := range msgs {
		var claim types.Claim
		err := json.Unmarshal(msg.Body, &claim)
		if err != nil {
			log.Printf("Error unmarshalling claim: %v", err)
			msg.Nack(false, true)
			continue
		}

		_, err = m.bankService.ProcessUnconfirmedClaim(&claim)
		if err != nil {
			log.Printf("Error processing claim: %v", err)
			msg.Nack(false, true)
		} else {
			log.Printf("Processed claim: %v", claim)
			msg.Ack(false)
		}
	}
}

func (m *Messaging) StartHotelOfferConsumer() {
	conn, channel, err := ConnectToRabbitMQ()
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer func(conn *amqp.Connection) {
		err := conn.Close()
		if err != nil {
			log.Printf("Error closing connection: %v", err)
		}
	}(conn)
	defer func(channel *amqp.Channel) {
		err := channel.Close()
		if err != nil {
			log.Printf("Error closing channel: %v", err)
		}
	}(channel)

	queue, err := channel.QueueDeclare(
		hotelOfferQueue,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to declare queue: %v", err)
	}

	msgs, err := channel.Consume(
		queue.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to register a consumer: %v", err)
	}

	log.Println("Waiting for hotel offers...")

	for msg := range msgs {
		var hotelOffer types.Room
		err := json.Unmarshal(msg.Body, &hotelOffer)
		if err != nil {
			log.Printf("Error unmarshalling hotel offer: %v", err)
			msg.Nack(false, true)
			continue
		}

		err = m.cacheHotelOfferInRedis(hotelOffer)
		if err != nil {
			log.Printf("Error caching hotel offer: %v", err)
			msg.Nack(false, true)
		} else {
			log.Printf("Successfully cached hotel offer: %v", hotelOffer)
			msg.Ack(false)
		}
	}
}

// cacheHotelOfferInRedis caches the hotel offer in Redis
func (m *Messaging) cacheHotelOfferInRedis(hotelOffer types.Room) error {
	ctx := context.Background()

	hotelOfferJSON, err := json.Marshal(hotelOffer)
	if err != nil {
		return fmt.Errorf("error marshalling hotel offer to JSON: %w", err)
	}

	err = m.redisClient.Set(ctx, fmt.Sprintf("hotel_offer:%s", hotelOffer.ID), hotelOfferJSON, redisCacheTimeout).Err()
	if err != nil {
		return fmt.Errorf("error caching hotel offer in Redis: %w", err)
	}

	log.Printf("Hotel offer with ID %s cached in Redis", hotelOffer.ID)
	return nil
}
