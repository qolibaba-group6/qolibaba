package app

import (
	"context"
	"errors"
	"qolibaba/config"
	"qolibaba/internal/user"
	userDomain "qolibaba/internal/user/domain"
	userPort "qolibaba/internal/user/port"
	"qolibaba/pkg/adapter/storage"
	"qolibaba/pkg/adapter/storage/types"
	"qolibaba/pkg/postgres"

	"gorm.io/gorm"
)

type app struct {
	db          *gorm.DB
	cfg         config.Config
	userService userPort.Service
}

func (a *app) DB() *gorm.DB {
	return a.db
}

func (a *app) Config() config.Config {
	return a.cfg
}

func (a *app) UserService(ctx context.Context) userPort.Service {
	return a.userService
}

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

	db = db.Debug()

	err = db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";").Error
	if err != nil {
		return err
	}

	err = db.AutoMigrate(
		&types.User{},
	)

	if err != nil {
		return err
	}

	err = a.setSuperAdmin(db)
	if err != nil {
		return err
	}

	a.db = db
	return nil
}

func (a *app) setSuperAdmin(db *gorm.DB) error {
	email := userDomain.Email(a.cfg.SuperAdmin.Email)
	if !email.IsValid() {
		return user.ErrInvalidEmail
	}

	err := db.Where("email = ?", string(email)).First(&userDomain.User{}).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		superAdmin := &types.User{
			Email:     a.cfg.SuperAdmin.Email,
			Password:  a.cfg.SuperAdmin.Password,
			IsAdmin:   true,
			Status:    uint8(userDomain.StatusActive),
			Role:      userDomain.RoleAdmin,
		}
		return db.Create(superAdmin).Error
	}

	return nil
}

func NewApp(cfg config.Config) (App, error) {
	a := &app{
		cfg: cfg,
	}
	
	if err := a.setDB(); err != nil {
		return nil, err
	}

	a.userService = user.NewService(storage.NewUserRepo(a.db))

	return a, nil
}

func NewMustApp(cfg config.Config) App {
	app, err := NewApp(cfg)
	if err != nil {
		panic(err)
	}
	return app
}
