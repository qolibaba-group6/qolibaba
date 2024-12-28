package hotel

import (
	"gorm.io/gorm"
	"qolibaba/config"
	"qolibaba/internal/hotels"
	"qolibaba/internal/hotels/port"
	"qolibaba/pkg/adapter/storage"
	"qolibaba/pkg/adapter/storage/types"
	"qolibaba/pkg/postgres"
)

type app struct {
	db           *gorm.DB
	cfg          config.Config
	hotelService port.Service
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

	if err := db.Exec("CREATE TYPE duration_type AS ENUM ('12 hours', '24 hours');").Error; err != nil {
		return err
	}
	if err := db.Exec("CREATE TYPE booking_status AS ENUM ('pending', 'confirmed', 'completed');").Error; err != nil {
		return err
	}

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

	// Initialize database.
	if err := a.setDB(); err != nil {
		return nil, err
	}

	// Initialize the hotel service.
	a.hotelService = hotels.NewService(storage.NewHotelRepo(a.db))

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
