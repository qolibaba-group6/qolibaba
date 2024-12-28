package service

import (
	"qolibaba/api/pb"
	"qolibaba/internal/routemap/domain"
)

func TerminalDomain2PB(t *domain.Terminal) *pb.Terminal {
	return &pb.Terminal{
		Id:           t.ID.String(),
		Name:         t.Name,
		TerminalType: uint32(t.Type),
		Country:      t.Country,
		State:        t.State,
		City:         t.City,
	}
}

func RouteDomain2PB(r *domain.Route) *pb.Route {
	return &pb.Route{
		Id:        r.ID.String(),
		RouteItem: &pb.RouteItem{
			Source:        TerminalDomain2PB(&r.Source),
			Destination:   TerminalDomain2PB(&r.Destination),
			RouteNumber:   uint32(r.RouteNumber),
			TransportType: uint32(r.TransportType),
			Distance:      float32(r.Distance),
		},
	}
}