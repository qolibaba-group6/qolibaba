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

func (s *routemapGRPCApi) GetTerminal(ctx context.Context, req *pb.TerminalGetByIDRequest) (*pb.Terminal, error) {
	return s.svc.GetTerminalByID(ctx, req)
}

func (s *routemapGRPCApi) CreateRoute(ctx context.Context, req *pb.CreateRouteRequest) (*pb.CreateRouteResponse, error) {
	return s.svc.CreateRoute(ctx, req)
}

func (s *routemapGRPCApi) GetRoute(ctx context.Context, req *pb.GetRouteByIDRequest) (*pb.Route, error) {
	return s.svc.GetRouteByID(ctx, req)
}

type routemapGRPCClient struct {
	cfg config.RoutemapServiceConfig
}

func NewRoutemapGRPCClient(cfg config.RoutemapServiceConfig) pb.RoutemapServiceClient {
	return &routemapGRPCClient{
		cfg: cfg,
	}
}

func (c *routemapGRPCClient) newClient() (pb.RoutemapServiceClient, error) {
	target := fmt.Sprintf("%s:%d", c.cfg.Host, c.cfg.Port)
	client, err := grpc.NewClient(target, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return pb.NewRoutemapServiceClient(client), nil
}

// CreateTerminal implements pb.RoutemapServiceClient.
func (c *routemapGRPCClient) CreateTerminal(ctx context.Context, in *pb.TerminalCreateRequest, opts ...grpc.CallOption) (*pb.TerminalCreateResponse, error) {
	client, err := c.newClient()
	if err != nil {
		return nil, err
	}

	res, err := client.CreateTerminal(ctx, in)

	if err != nil {
		return nil, err
	}

	return  res, nil
}

func (c *routemapGRPCClient) GetTerminal(ctx context.Context, in *pb.TerminalGetByIDRequest, opts ...grpc.CallOption) (*pb.Terminal, error) {
	client, err := c.newClient()
	if err != nil {
		return nil, err
	}

	res, err := client.GetTerminal(ctx, in)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c *routemapGRPCClient) CreateRoute(ctx context.Context, in *pb.CreateRouteRequest, opts ...grpc.CallOption) (*pb.CreateRouteResponse, error) {
	client, err := c.newClient()
	if err != nil {
		return nil, err
	}

	resp, err := client.CreateRoute(ctx, in)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *routemapGRPCClient) GetRoute(ctx context.Context, in *pb.GetRouteByIDRequest, opts ...grpc.CallOption) (*pb.Route, error) {
	client, err := c.newClient()
	if err != nil {
		return nil, err
	}

	resp, err := client.GetRoute(ctx, in)
	if err != nil {
		return nil, err
	}

	return resp, nil
}