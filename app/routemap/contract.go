package routemap

import (
	"qolibaba/config"
	"qolibaba/internal/routemap/port"

	"gorm.io/gorm"
)

type App interface {
	DB() *gorm.DB
	Config() config.Config
	RoutemapService() port.Service
}