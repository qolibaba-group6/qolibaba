package app

import (
	"context"
	"qolibaba/config"
	userPort "qolibaba/internal/user/port"

	"gorm.io/gorm"
)

type App interface {
	UserService(ctx context.Context) userPort.Service
	DB() *gorm.DB
	Config() config.Config
}
