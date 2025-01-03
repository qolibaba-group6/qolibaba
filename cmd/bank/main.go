package main

import (
	"flag"
	"github.com/streadway/amqp"
	"log"
	"os"
	"qolibaba/api/handlers/http"
	"qolibaba/app/bank"
	"qolibaba/config"
	"qolibaba/pkg/messaging"
)

var (
	configPath = flag.String("config", "config.json", "service configuration file")
)

func main() {
	flag.Parse()

	if v := os.Getenv("CONFIG_PATH"); len(v) > 0 {
		*configPath = v
	}

	cfg := config.MustReadConfig(*configPath)

	bankApp, err := bank.NewApp(cfg)
	if err != nil {
		log.Fatalf("failed to initialize bank app: %v", err)
	}

	rabbitMQConn, rabbitMQChannel, err := messaging.ConnectToRabbitMQ()
	if err != nil {
		log.Fatalf("failed to connect to RabbitMQ: %v", err)
	}
	defer func(rabbitMQConn *amqp.Connection) {
		err := rabbitMQConn.Close()
		if err != nil {

		}
	}(rabbitMQConn)
	defer func(rabbitMQChannel *amqp.Channel) {
		err := rabbitMQChannel.Close()
		if err != nil {

		}
	}(rabbitMQChannel)

	rabbitMQConsumer := messaging.NewMessaging(rabbitMQChannel, bankApp.BankService())

	go rabbitMQConsumer.StartClaimConsumer()

	err = http.RunBank(bankApp, cfg.Server)
	if err != nil {
		log.Fatalf("failed to start HTTP server: %v", err)
	}
}
