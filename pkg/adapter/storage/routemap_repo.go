package storage

import (
	"context"
	"errors"
	routemapDomain "qolibaba/internal/routemap/domain"
	routemapPort "qolibaba/internal/routemap/port"
	"qolibaba/pkg/adapter/storage/mapper"
	"qolibaba/pkg/adapter/storage/types"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type routemapRepo struct {
	db *gorm.DB
}

func NewRouteMapRepo(db *gorm.DB) routemapPort.Repo {
	return &routemapRepo{
		db: db,
	}
}

func (r *routemapRepo) CreateTerminal(ctx context.Context, terminal routemapDomain.Terminal) (routemapDomain.TerminalUUID, error) {
	t := mapper.TerminalDomain2Storage(terminal)
	return t.ID, r.db.WithContext(ctx).Table("terminals").Create(t).Error
}

func (r *routemapRepo) CreateRoute(ctx context.Context, route routemapDomain.Route) (routemapDomain.RouteUUID, error) {
	panic("not implemented")
}

func (r *routemapRepo) GetTerminalByID(ctx context.Context, id routemapDomain.TerminalUUID) (*routemapDomain.Terminal, error) {
	var terminal types.Terminal

	q := r.db.Table("terminals").Debug().WithContext(ctx)

	err := q.Where("id = ?", id).First(&terminal).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	if terminal.ID == uuid.Nil {
		return nil, nil
	}

	return mapper.TerminalStorage2Domain(terminal), nil
}

func (r *routemapRepo) GetRouteByID(ctx context.Context, id routemapDomain.RouteUUID) (*routemapDomain.Route, error) {
	panic("not implemented")
}

func (r *routemapRepo) GetTerminal(ctx context.Context, filter routemapDomain.TerminalFilter) ([]routemapDomain.Terminal, error) {
	panic("not implemented")
}

func (r *routemapRepo) GetRoute(ctx context.Context, filter routemapDomain.RouteFilter) ([]routemapDomain.Route, error) {
	panic("not implemented")
}
