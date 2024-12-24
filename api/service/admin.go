package service

import (
	"context"
	"qolibaba/api/pb"
	"qolibaba/internal/admin/port"
)


type AdminService struct {
	svc port.Service
}

func NewAdminService(svc port.Service) *AdminService {
	return &AdminService{
		svc: svc,
	}
}

func (s *AdminService) SayHello(ctx context.Context, req *pb.AdminSayHelloRequest) (*pb.AdminSayHelloResponse ,error) {
	adminSays, err := s.svc.SayHello(ctx, req.GetName())
	if err != nil {
		return nil, err
	}
	
	return &pb.AdminSayHelloResponse{
		AdminSays: adminSays,
	}, nil
}