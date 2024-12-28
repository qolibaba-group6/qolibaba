package service

import (
	"context"
	"qolibaba/api/pb"
	"qolibaba/internal/routemap"
	"qolibaba/internal/routemap/domain"
	"qolibaba/internal/routemap/port"

	"github.com/google/uuid"
)

var (
	ErrTerminalOnCreate = routemap.ErrTerminalOnCreate
)

type RoutemapService struct {
	svc port.Service
}

func NewRoutemapService(svc port.Service) *RoutemapService {
	return &RoutemapService{
		svc: svc,
	}
}

func (s *RoutemapService) CreateTerminal(ctx context.Context, req *pb.TerminalCreateRequest) (*pb.TerminalCreateResponse, error) {
	id, err := s.svc.CreateTerminal(ctx, domain.Terminal{
		ID:      [16]byte{},
		Name:    req.GetName(),
		Type:    domain.TerminalType(req.GetTerminalType()),
		Country: req.GetCountry(),
		State:   req.GetState(),
		City:    req.GetCity(),
	})

	if err != nil {
		return nil, ErrTerminalOnCreate
	}

	return &pb.TerminalCreateResponse{
		TerminalID: id.String(),
	}, err
}

func (s *RoutemapService) GetTErminalByID(ctx context.Context, req *pb.TerminalGetByIDRequest) (*pb.Terminal, error) {
	terminalID, err := uuid.Parse(req.GetTerminalID())
	if err != nil {
		return nil, err
	}

	terminal, err := s.svc.GetTerminalByID(ctx, terminalID)
	if err != nil {
		return nil, err
	}

	return &pb.Terminal{
		Id:           terminal.ID.String(),
		Name:         terminal.Name,
		TerminalType: uint32(terminal.Type),
		Country:      terminal.Country,
		State:        terminal.State,
		City:         terminal.City,
	}, nil
}
