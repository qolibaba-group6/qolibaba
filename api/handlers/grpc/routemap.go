package grpc

import (
	"context"
	"fmt"
	"qolibaba/api/pb"
	"qolibaba/api/service"
	"qolibaba/config"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type routemapGRPCApi struct {
	pb.UnimplementedRoutemapServiceServer
	svc service.RoutemapService
}

func NewRoutemapGRPCServer(svc service.RoutemapService) pb.RoutemapServiceServer {
	return &routemapGRPCApi{
		svc: svc,
	}
}

func (s *routemapGRPCApi) CreateTerminal(ctx context.Context, req *pb.TerminalCreateRequest) (*pb.TerminalCreateResponse, error) {
	return s.svc.CreateTerminal(ctx, req)
}

type routemapGRPCClient struct {
	cfg config.RoutemapServiceConfig
}

// CreateTerminal implements pb.RoutemapServiceClient.
func (c *routemapGRPCClient) CreateTerminal(ctx context.Context, in *pb.TerminalCreateRequest, opts ...grpc.CallOption) (*pb.TerminalCreateResponse, error) {
	target := fmt.Sprintf(":%d", c.cfg.Port)
	client, err := grpc.NewClient(target, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	serviceClient := pb.NewRoutemapServiceClient(client)

	res, err := serviceClient.CreateTerminal(ctx, in)

	if err != nil {
		return nil, err
	}

	return  res, nil
}

func NewRoutemapGRPCClient(cfg config.RoutemapServiceConfig) pb.RoutemapServiceClient {
	return &routemapGRPCClient{
		cfg: cfg,
	}
}
