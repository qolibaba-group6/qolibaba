package bank

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"log"
	"qolibaba/config"
	"qolibaba/internal/bank"
	"qolibaba/internal/bank/port"
	agenciesPort "qolibaba/internal/travel_agencies/port"
	"qolibaba/pkg/adapter/storage"
	"qolibaba/pkg/adapter/storage/types"
	"qolibaba/pkg/postgres"
)

type app struct {
	db              *gorm.DB
	cfg             config.Config
	bankService     port.Service
	agenciesService agenciesPort.Service
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

// BankService provides access to the bank service implementation.
func (a *app) BankService() port.Service {
	return a.bankService
}

// setRedis initializes the Redis connection.
func (a *app) setRedis() error {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", a.cfg.Redis.Host, a.cfg.Redis.Port),
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
	/*if err := db.Exec("CREATE TYPE wallet_role AS ENUM ('user', 'bank');").Error; err != nil {
		return err
	}
	if err := db.Exec("CREATE TYPE transaction_type AS ENUM ('deposit', 'withdrawal');").Error; err != nil {
		return err
	}
	if err := db.Exec("CREATE TYPE status AS ENUM ('pending', 'paid', 'completed', 'failed');").Error; err != nil {
		return err
	}*/

	if err := db.AutoMigrate(
		&types.Wallet{},
		&types.Transaction{},
		&types.Claim{},
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

	if err := a.setDB(); err != nil {
		log.Printf("Error initializing database: %v", err)
		return nil, err
	}
	a.bankService = bank.NewService(storage.NewBankRepo(a.db))

	return a, nil
}

// NewMustApp initializes a new app instance and panics if an error occurs.
func NewMustApp(cfg config.Config) App {
	app, err := NewApp(cfg)
	if err != nil {
		log.Panicf("Error initializing app: %v", err)
	}
	return app
}
