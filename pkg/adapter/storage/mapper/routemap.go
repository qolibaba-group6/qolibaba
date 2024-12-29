package mapper

import (
	"qolibaba/internal/routemap/domain"
	"qolibaba/pkg/adapter/storage/types"
)


func TerminalDomain2Storage(t domain.Terminal) *types.Terminal {
	return &types.Terminal{
		Model:   types.Model{
			ID: t.ID,
		},
		Name:    t.Name,
		Type:    uint8(t.Type),
		Country: t.Country,
		State:   t.State,
		City:    t.City,
	}
}

func TerminalStorage2Domain(t types.Terminal) *domain.Terminal {
	return &domain.Terminal{
		ID:      t.ID,
		Name:    t.Name,
		Type:    domain.TerminalType(t.Type),
		Country: t.Country,
		State:   t.State,
		City:    t.City,
	}
}

func RouteDomain2Storage(r domain.Route) *types.Route {
	return &types.Route{
		Model:         types.Model{
			ID:        r.ID,
		},
		SourceID:      r.Source.ID,
		DestinationID: r.Destination.ID,
		RouteNumber:   r.RouteNumber,
		TransportType: uint8(r.TransportType),
		Distance:      r.Distance,
	}
}

func RouteStorage2Domain(r types.Route) *domain.Route {
	return &domain.Route{
		ID:            r.ID,
		Source:        domain.Terminal{
			ID:      r.SourceID,
		},
		Destination:   domain.Terminal{
			ID: r.DestinationID,
		},
		RouteNumber:   r.RouteNumber,
		TransportType: domain.TransportType(r.TransportType),
		Distance:      r.Distance,
	}
}

