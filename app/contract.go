package app

import (
	"qolibaba/config"

	"gorm.io/gorm"
)

type App interface {
	DB() *gorm.DB
	Config() config.Config
}