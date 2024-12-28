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

func (s *RoutemapService) CreateRoute(ctx context.Context, req *pb.CreateRouteRequest) (*pb.CreateRouteResponse, error) {
	sourceId, err := uuid.Parse(req.RouteItem.Source.Id)
	if err != nil {
		return nil, err
	}

	destinationId, err := uuid.Parse(req.RouteItem.Destination.Id)
	if err != nil {
		return nil, err
	}

	id, err := s.svc.CreateRoute(ctx, domain.Route{
		Source: domain.Terminal{
			ID: sourceId,
		},
		Destination: domain.Terminal{
			ID: destinationId,
		},
		RouteNumber:   uint(req.RouteItem.RouteNumber),
		TransportType: domain.TransportType(req.RouteItem.TransportType),
		Distance:      float64(req.RouteItem.Distance),
	})

	if err != nil {
		return nil, err
	}

	return &pb.CreateRouteResponse{
		Id: id.String(),
	}, nil
}
