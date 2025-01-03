package travel_agency

import (
	"gorm.io/gorm"
	"qolibaba/config"
	"qolibaba/internal/travel_agencies/port"
)

type App interface {
	DB() *gorm.DB
	Config() config.Config
	TravelAgency() port.Service
}
