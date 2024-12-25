package admin

import (
	"qolibaba/config"
	"qolibaba/internal/admin"
	"qolibaba/internal/admin/port"
	"qolibaba/pkg/adapter/storage"
	"qolibaba/pkg/postgres"

	"gorm.io/gorm"
)


type app struct {
	db          *gorm.DB
	cfg         config.Config
	adminService port.Service
}

func (a *app) DB() *gorm.DB {
	return a.db
}

func (a *app) Config() config.Config {
	return a.cfg
}

func (a *app) AdminService() port.Service {
	return a.adminService
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

	// db = db.Debug()

	err = db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";").Error
	if err != nil {
		return err
	}

	err = db.AutoMigrate(
		//
	)
	if err != nil {
		return err
	}

	a.db = db
	return nil
}


func NewApp(cfg config.Config) (App, error) {
	a := &app{
		cfg: cfg,
	}
	
	if err := a.setDB(); err != nil {
		return nil, err
	}

	a.adminService = admin.NewService(storage.NewAdminRepo(a.db))

	return a, nil
}

func NewMustApp(cfg config.Config) App {
	app, err := NewApp(cfg)
	if err != nil {
		panic(err)
	}
	return app
}
