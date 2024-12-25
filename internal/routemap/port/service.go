package port

import (
	"context"
	"qolibaba/internal/routemap/domain"
)

type Service interface {
	CreateTerminal(ctx context.Context, terminal domain.Terminal) (domain.TerminalUUID, error)
	GetTerminalByID(ctx context.Context, terminalID domain.TerminalUUID) (*domain.Terminal, error)
	CreateRoute(ctx context.Context, route domain.Route) (domain.RouteUUID, error)
	GetRouteByID(ctx context.Context, routeID domain.RouteUUID) (*domain.Route, error)
}
