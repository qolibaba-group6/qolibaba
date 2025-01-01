package grpc

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/ehsansobhani/travel_agencies/api/pb"
	"github.com/ehsansobhani/travel_agencies/internal/travel_agencies/service"
	"google.golang.org/protobuf/types/known/emptypb"
)

type TravelAgenciesGRPCServer struct {
	pb.UnimplementedTravelAgenciesServiceServer
	CompanyService *service.CompanyService
	TripService    *service.TripService
}

func NewTravelAgenciesGRPCServer(companyService *service.CompanyService, tripService *service.TripService) *TravelAgenciesGRPCServer {
	return &TravelAgenciesGRPCServer{
		CompanyService: companyService,
		TripService:    tripService,
	}
}

// Company Methods

func (s *TravelAgenciesGRPCServer) CreateCompany(ctx context.Context, req *pb.CreateCompanyRequest) (*pb.Company, error) {
	company, err := s.CompanyService.CreateCompany(req.Name, req.Owner)
	if err != nil {
		return nil, err
	}
	return &pb.Company{
		Id:         company.ID.String(),
		Name:       company.Name,
		Owner:      company.Owner,
		Created_at: company.CreatedAt.Format(time.RFC3339),
		Updated_at: company.UpdatedAt.Format(time.RFC3339),
	}, nil
}

func (s *TravelAgenciesGRPCServer) GetCompany(ctx context.Context, req *pb.GetCompanyRequest) (*pb.Company, error) {
	id, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, err
	}
	company, err := s.CompanyService.GetCompany(id)
	if err != nil {
		return nil, err
	}
	return &pb.Company{
		Id:         company.ID.String(),
		Name:       company.Name,
		Owner:      company.Owner,
		Created_at: company.CreatedAt.Format(time.RFC3339),
		Updated_at: company.UpdatedAt.Format(time.RFC3339),
	}, nil
}

func (s *TravelAgenciesGRPCServer) UpdateCompany(ctx context.Context, req *pb.UpdateCompanyRequest) (*pb.Company, error) {
	id, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, err
	}
	company, err := s.CompanyService.UpdateCompany(id, req.Name, req.Owner)
	if err != nil {
		return nil, err
	}
	return &pb.Company{
		Id:         company.ID.String(),
		Name:       company.Name,
		Owner:      company.Owner,
		Created_at: company.CreatedAt.Format(time.RFC3339),
		Updated_at: company.UpdatedAt.Format(time.RFC3339),
	}, nil
}

func (s *TravelAgenciesGRPCServer) DeleteCompany(ctx context.Context, req *pb.DeleteCompanyRequest) (*emptypb.Empty, error) {
	id, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, err
	}
	err = s.CompanyService.DeleteCompany(id)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (s *TravelAgenciesGRPCServer) ListCompanies(ctx context.Context, req *pb.ListCompaniesRequest) (*pb.ListCompaniesResponse, error) {
	companies, err := s.CompanyService.ListCompanies()
	if err != nil {
		return nil, err
	}
	var pbCompanies []*pb.Company
	for _, c := range companies {
		pbCompanies = append(pbCompanies, &pb.Company{
			Id:         c.ID.String(),
			Name:       c.Name,
			Owner:      c.Owner,
			Created_at: c.CreatedAt.Format(time.RFC3339),
			Updated_at: c.UpdatedAt.Format(time.RFC3339),
		})
	}
	return &pb.ListCompaniesResponse{
		Companies: pbCompanies,
	}, nil
}

// Trip Methods

func (s *TravelAgenciesGRPCServer) CreateTrip(ctx context.Context, req *pb.CreateTripRequest) (*pb.Trip, error) {
	trip, err := s.TripService.CreateTrip(req.Company_id, req.Type, req.Origin, req.Destination, req.Departure_time, req.Release_date, req.Tariff, req.Status)
	if err != nil {
		return nil, err
	}
	return &pb.Trip{
		Id:             trip.ID.String(),
		Company_id:     trip.CompanyID.String(),
		Type:           trip.Type,
		Origin:         trip.Origin,
		Destination:    trip.Destination,
		Departure_time: trip.DepartureTime.Format(time.RFC3339),
		Release_date:   trip.ReleaseDate.Format(time.RFC3339),
		Tariff:         trip.Tariff,
		Status:         trip.Status,
		Created_at:     trip.CreatedAt.Format(time.RFC3339),
		Updated_at:     trip.UpdatedAt.Format(time.RFC3339),
	}, nil
}

func (s *TravelAgenciesGRPCServer) GetTrip(ctx context.Context, req *pb.GetTripRequest) (*pb.Trip, error) {
	id, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, err
	}
	trip, err := s.TripService.GetTrip(id)
	if err != nil {
		return nil, err
	}
	return &pb.Trip{
		Id:             trip.ID.String(),
		Company_id:     trip.CompanyID.String(),
		Type:           trip.Type,
		Origin:         trip.Origin,
		Destination:    trip.Destination,
		Departure_time: trip.DepartureTime.Format(time.RFC3339),
		Release_date:   trip.ReleaseDate.Format(time.RFC3339),
		Tariff:         trip.Tariff,
		Status:         trip.Status,
		Created_at:     trip.CreatedAt.Format(time.RFC3339),
		Updated_at:     trip.UpdatedAt.Format(time.RFC3339),
	}, nil
}

func (s *TravelAgenciesGRPCServer) UpdateTrip(ctx context.Context, req *pb.UpdateTripRequest) (*pb.Trip, error) {
	id, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, err
	}
	companyID, err := uuid.Parse(req.Company_id)
	if err != nil {
		return nil, err
	}
	trip, err := s.TripService.UpdateTrip(id, companyID, req.Type, req.Origin, req.Destination, req.Departure_time, req.Release_date, req.Tariff, req.Status)
	if err != nil {
		return nil, err
	}
	return &pb.Trip{
		Id:             trip.ID.String(),
		Company_id:     trip.CompanyID.String(),
		Type:           trip.Type,
		Origin:         trip.Origin,
		Destination:    trip.Destination,
		Departure_time: trip.DepartureTime.Format(time.RFC3339),
		Release_date:   trip.ReleaseDate.Format(time.RFC3339),
		Tariff:         trip.Tariff,
		Status:         trip.Status,
		Created_at:     trip.CreatedAt.Format(time.RFC3339),
		Updated_at:     trip.UpdatedAt.Format(time.RFC3339),
	}, nil
}

func (s *TravelAgenciesGRPCServer) DeleteTrip(ctx context.Context, req *pb.DeleteTripRequest) (*emptypb.Empty, error) {
	id, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, err
	}
	err = s.TripService.DeleteTrip(id)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (s *TravelAgenciesGRPCServer) ListTrips(ctx context.Context, req *pb.ListTripsRequest) (*pb.ListTripsResponse, error) {
	trips, err := s.TripService.ListTrips(req.Company_id)
	if err != nil {
		return nil, err
	}
	var pbTrips []*pb.Trip
	for _, t := range trips {
		pbTrips = append(pbTrips, &pb.Trip{
			Id:             t.ID.String(),
			Company_id:     t.CompanyID.String(),
			Type:           t.Type,
			Origin:         t.Origin,
			Destination:    t.Destination,
			Departure_time: t.DepartureTime.Format(time.RFC3339),
			Release_date:   t.ReleaseDate.Format(time.RFC3339),
			Tariff:         t.Tariff,
			Status:         t.Status,
			Created_at:     t.CreatedAt.Format(time.RFC3339),
			Updated_at:     t.UpdatedAt.Format(time.RFC3339),
		})
	}
	return &pb.ListTripsResponse{
		Trips: pbTrips,
	}, nil
}
