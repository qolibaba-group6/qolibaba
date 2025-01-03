package storage

import (
	adminPort "qolibaba/internal/admin/port"

	"gorm.io/gorm"
)

type adminRepo struct {
	db *gorm.DB
}

func NewAdminRepo(db *gorm.DB) adminPort.Repo {
	return &adminRepo{
		db: db,
	}
}