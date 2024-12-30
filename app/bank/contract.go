package bank

import (
	"gorm.io/gorm"
	"qolibaba/config"
	"qolibaba/internal/bank/port"
)

type App interface {
	DB() *gorm.DB
	Config() config.Config
	BankService() port.Service
}
