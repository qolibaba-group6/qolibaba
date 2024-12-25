package routemap

import (
	"context"
	"qolibaba/internal/routemap/domain"
	"qolibaba/internal/routemap/port"
)

type service struct {
	repo port.Repo
}

func NewService(repo port.Repo) port.Service {
	return &service{
		repo: repo,
	}
}

func (s *service) CreateTerminal(ctx context.Context, terminal domain.Terminal) (domain.TerminalUUID, error) {
	panic("not implemented")
}

func (s *service) GetTerminalByID(ctx context.Context, terminalID domain.TerminalUUID) (*domain.Terminal, error) {
	panic("not implemented")
}

func (s *service) CreateRoute(ctx context.Context, route domain.Route) (domain.RouteUUID, error) {
	panic("not implemented")
}

func (s *service) GetRouteByID(ctx context.Context, routeID domain.RouteUUID) (*domain.Route, error) {
	panic("not implemented")
}
