// pkg/postgres/gorm.go
package postgres

import (
	"github.com/ehsansobhani/travel_agencies/internal/travel_agencies/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// GormDB wraps the GORM DB instance
type GormDB struct {
	*gorm.DB
}

// NewGormDB initializes a new GORM database connection
func NewGormDB(dsn string) (*GormDB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, err
	}

	// اتوماتیک مهاجرت مدل‌های دامنه
	err = db.AutoMigrate(&domain.Company{}, &domain.Trip{})
	if err != nil {
		return nil, err
	}

	return &GormDB{DB: db}, nil
}
