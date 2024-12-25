package service

import (
	"context"
	"qolibaba/api/pb"
	"qolibaba/config"
	"qolibaba/internal/admin/port"
)


type AdminService struct {
	svc port.Service
	cfg config.AdminServiceConfig
}

func NewAdminService(svc port.Service, cfg config.AdminServiceConfig) *AdminService {
	return &AdminService{
		svc: svc,
		cfg: cfg,
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