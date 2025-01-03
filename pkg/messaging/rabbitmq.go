package messaging

import (
	"encoding/json"
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"qolibaba/internal/bank/port"
	"qolibaba/pkg/adapter/storage/types"
)

const (
	rabbitmqURL = "amqp://guest:guest@localhost:5672/"
	claimsQueue = "claimsQueue"
)

type Messaging struct {
	channel     *amqp.Channel
	bankService port.Service
}

func NewMessaging(channel *amqp.Channel, bankService port.Service) *Messaging {
	return &Messaging{
		channel:     channel,
		bankService: bankService,
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

func (m *Messaging) PublishClaimToBank(claimData []byte) (string, error) {
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
		return "", fmt.Errorf("failed to publish claim: %w", err)
	}

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
		return "", fmt.Errorf("failed to consume response: %w", err)
	}

	for response := range msg {
		var claimResp struct {
			ClaimID string `json:"claim_id"`
		}
		err := json.Unmarshal(response.Body, &claimResp)
		if err != nil {
			return "", fmt.Errorf("error unmarshalling claim response: %w", err)
		}
		return claimResp.ClaimID, nil
	}

	return "", fmt.Errorf("no response from Bank Service")
}

func (m *Messaging) StartClaimConsumer() {
	conn, channel, err := ConnectToRabbitMQ()
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer func(conn *amqp.Connection) {
		err := conn.Close()
		if err != nil {

		}
	}(conn)
	defer func(channel *amqp.Channel) {
		err := channel.Close()
		if err != nil {

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
		true,
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
			continue
		}

		_, err = m.bankService.ProcessUnconfirmedClaim(&claim)
		if err != nil {
			log.Printf("Error processing claim: %v", err)
		} else {
			log.Printf("Processed claim: %v", claim)
		}
	}
}
