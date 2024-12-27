package routemap

import (
	"context"
	"errors"
	"qolibaba/internal/routemap/domain"
	"qolibaba/internal/routemap/port"
)

var (
	ErrTerminalOnCreate    = errors.New("error on creating new terminal")
	ErrTerminalInvalidType = errors.New("invalid terminal type")
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
	if ok := terminal.Type.IsValid(); !ok {
		return domain.NilUUID(), ErrTerminalInvalidType
	}

	terminalID, err := s.repo.CreateTerminal(ctx, terminal)
	if err != nil {
		// log
		return domain.NilUUID(), ErrTerminalOnCreate
	}

	return terminalID, nil
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

func (s *service) GetTerminal(ctx context.Context, filter domain.TerminalFilter) ([]domain.Terminal, error) {
	panic("not implemented")
}

func (s *service) GetRoute(ctx context.Context, filter domain.RouteFilter) ([]domain.Route, error) {
	panic("not implemented")
}
