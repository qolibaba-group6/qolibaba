package port

import (
	"context"
	"qolibaba/internal/routemap/domain"
)

type Repo interface {
	CreateTerminal(ctx context.Context, terminal domain.Terminal) (domain.TerminalUUID, error)
	CreateRoute(ctx context.Context, route domain.Route) (domain.RouteUUID, error)
	GetTerminalByID(ctx context.Context, id domain.TerminalUUID) (*domain.Terminal, error)
	GetRouteByID(ctx context.Context, id domain.RouteUUID) (*domain.Route, error)
	GetTerminal(ctx context.Context, filter domain.TerminalFilter) ([]domain.Terminal, error)
	GetRoute(ctx context.Context, filter domain.RouteFilter) ([]domain.Route, error)
}
