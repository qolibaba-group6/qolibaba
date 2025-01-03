package admin

import (
	"qolibaba/config"
	"qolibaba/internal/admin/port"

	"gorm.io/gorm"
)

type App interface {
	DB() *gorm.DB
	Config() config.Config
	AdminService() port.Service
}
