package app

import (
	"context"
	"qolibaba/app/hotel"
	"qolibaba/config"
	userPort "qolibaba/internal/user/port"

	"gorm.io/gorm"
)

type App interface {
	UserService(ctx context.Context) userPort.Service
	HotelService(ctx context.Context) hotel.App
	DB() *gorm.DB
	Config() config.Config
}
