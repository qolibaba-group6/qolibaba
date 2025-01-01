package hotel

import (
	"context"
	"github.com/gofiber/fiber/v2/log"
	"github.com/redis/go-redis/v9"
	"github.com/streadway/amqp"
	"gorm.io/gorm"
	"qolibaba/config"
	bankPort "qolibaba/internal/bank/port"
	"qolibaba/internal/hotels"
	"qolibaba/internal/hotels/port"
	agenciesPort "qolibaba/internal/travel_agencies/port"
	"qolibaba/pkg/adapter/storage"
	"qolibaba/pkg/adapter/storage/types"
	"qolibaba/pkg/messaging"
	"qolibaba/pkg/postgres"
)

type app struct {
	db              *gorm.DB
	cfg             config.Config
	bankService     bankPort.Service
	agenciesService agenciesPort.Service
	hotelService    port.Service
	redisClient     *redis.Client
}

// DB provides access to the database instance.
func (a *app) DB() *gorm.DB {
	return a.db
}

// Config provides access to the application configuration.
func (a *app) Config() config.Config {
	return a.cfg
}

// HotelService provides access to the hotel service implementation.
func (a *app) HotelService() port.Service {
	return a.hotelService
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

	// Ensure required extensions are available.
	if err := db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";").Error; err != nil {
		return err
	}

	// Create custom enum types if they don't exist
	/*
		if err := db.Exec("CREATE TYPE room_status AS ENUM ('free', 'booked');").Error; err != nil {
			return err
		}
		if err := db.Exec("CREATE TYPE duration_type AS ENUM ('12 hours', '24 hours');").Error; err != nil {
			return err
		}
		if err := db.Exec("CREATE TYPE booking_status AS ENUM ('pending', 'confirmed', 'completed');").Error; err != nil {
			return err
		}
	*/
	// Apply database migrations for hotel-related models.
	if err := db.AutoMigrate(
		&types.Hotel{},
		&types.Room{},
		&types.Booking{},
	); err != nil {
		return err
	}

	a.db = db
	return nil
}

// NewApp initializes and returns a new app instance with the provided configuration.
func NewApp(cfg config.Config) (App, error) {
	a := &app{
		cfg: cfg,
	}

	// Initialize the database connection
	if err := a.setDB(); err != nil {
		return nil, err
	}

	// Initialize Redis if needed
	if err := a.setRedis(); err != nil {
		return nil, err
	}

	// Connect to RabbitMQ
	conn, channel, err := messaging.ConnectToRabbitMQ()
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

	messagingClient := messaging.NewMessaging(channel, a.bankService, a.redisClient, a.agenciesService)

	a.hotelService = hotels.NewService(storage.NewHotelRepo(a.db), messagingClient)

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
