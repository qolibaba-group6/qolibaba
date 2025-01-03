package travel_agency

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/streadway/amqp"
	"gorm.io/gorm"
	"log"
	"qolibaba/config"
	bankPort "qolibaba/internal/bank/port"
	"qolibaba/internal/travel_agencies"
	travelPort "qolibaba/internal/travel_agencies/port"
	"qolibaba/pkg/adapter/storage"
	"qolibaba/pkg/adapter/storage/types"
	"qolibaba/pkg/messaging"
	"qolibaba/pkg/postgres"
)

type app struct {
	db            *gorm.DB
	cfg           config.Config
	travelService travelPort.Service
	bankService   bankPort.Service
	redisClient   *redis.Client
}

// DB provides access to the database instance.
func (a *app) DB() *gorm.DB {
	return a.db
}

// Config provides access to the application configuration.
func (a *app) Config() config.Config {
	return a.cfg
}

// TravelAgency provides access to the travel agency service implementation.
func (a *app) TravelAgency() travelPort.Service {
	return a.travelService
}

// RedisClient provides access to the Redis client instance.
func (a *app) RedisClient() *redis.Client {
	return a.redisClient
}

// setDB initializes the database connection and applies migrations.
func (a *app) setDB() error {
	db, err := postgres.NewPsqlGormConnection(postgres.DBConnOptions{
		User:   a.cfg.DB.User,
		Pass:   a.cfg.DB.Password,
		Host:   a.cfg.DB.Host,
		Port:   a.cfg.DB.Port,
		DBName: a.cfg.DB.Database,
		Schema: a.cfg.DB.Schema,
	})
	if err != nil {
		return err
	}

	// Apply database migrations for travel agency-related models.
	if err := db.AutoMigrate(
		&types.Tour{},
		&types.TravelAgency{},
		&types.TourBooking{},
	); err != nil {
		return err
	}

	a.db = db
	return nil
}

// setRedis initializes the Redis connection.
func (a *app) setRedis() error {
	rdb := redis.NewClient(&redis.Options{
		Addr:     a.cfg.Redis.Host,
		Password: a.cfg.Redis.Password,
		DB:       0,
	})

	ctx := context.Background()
	if _, err := rdb.Ping(ctx).Result(); err != nil {
		return err
	}

	a.redisClient = rdb
	return nil
}

// NewApp initializes and returns a new app instance with the provided configuration.
func NewApp(cfg config.Config) (App, error) {
	a := &app{
		cfg: cfg,
	}

	// Connect to RabbitMQ
	conn, channel, err := messaging.ConnectToRabbitMQ()
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer func(conn *amqp.Connection) {
		if err := conn.Close(); err != nil {
			log.Printf("Error closing RabbitMQ connection: %v", err)
		}
	}(conn)

	defer func(channel *amqp.Channel) {
		if err := channel.Close(); err != nil {
			log.Printf("Error closing RabbitMQ channel: %v", err)
		}
	}(channel)

	messagingClient := messaging.NewMessaging(channel, a.redisClient)

	go func() {
		err := messagingClient.StartConsumer(messaging.TourQueue, messagingClient.HandleClaim)
		if err != nil {

		}
	}()
	// Set up database connection
	if err := a.setDB(); err != nil {
		return nil, fmt.Errorf("failed to set up database: %v", err)
	}

	// Initialize Redis client
	if err := a.setRedis(); err != nil {
		return nil, fmt.Errorf("failed to set up Redis: %v", err)
	}
	a.travelService = travel_agencies.NewTravelAgencyService(
		storage.NewTravelAgencyRepository(a.db),
		messagingClient,
		a.redisClient,
	)

	return a, nil
}

// NewMustApp initializes a new app instance and panics if an error occurs.
func NewMustApp(cfg config.Config) App {
	app, err := NewApp(cfg)
	if err != nil {
		panic(err)
	}
	return app
}
