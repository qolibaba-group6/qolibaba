package storage

import (
	"context"
	routemapDomain "qolibaba/internal/routemap/domain"
	routemapPort "qolibaba/internal/routemap/port"

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
	panic("not implemented")
}
func (r *routemapRepo) CreateRoute(ctx context.Context, route routemapDomain.Route) (routemapDomain.RouteUUID, error) {
	panic("not implemented")
}
func (r *routemapRepo) GetTerminalByID(ctx context.Context, id routemapDomain.TerminalUUID) (*routemapDomain.Terminal, error) {
	panic("not implemented")
}
func (r *routemapRepo) GetRouteByID(ctx context.Context, id routemapDomain.RouteUUID) (*routemapDomain.Route, error) {
	panic("not implemented")
}
