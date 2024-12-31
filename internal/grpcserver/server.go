package grpcserver

import (
	"context"
	"log"
	"net"

	pb "final12/proto"

	"google.golang.org/grpc"
)

type CompanyServiceServer struct {
	pb.UnimplementedCompanyServiceServer
}

func NewCompanyServiceServer() *CompanyServiceServer {
	return &CompanyServiceServer{}
}

func (s *CompanyServiceServer) GetCompany(ctx context.Context, req *pb.CompanyRequest) (*pb.CompanyResponse, error) {
	log.Printf("Received request for company ID: %s", req.Id)
	// Simulate fetching company details
	return &pb.CompanyResponse{
		Id:    req.Id,
		Name:  "Example Company",
		Owner: "Example Owner",
	}, nil
}

func (s *CompanyServiceServer) CreateCompany(ctx context.Context, req *pb.NewCompany) (*pb.CompanyResponse, error) {
	log.Printf("Creating company: %s owned by %s", req.Name, req.Owner)
	// Simulate creating a company
	return &pb.CompanyResponse{
		Id:    "new-id",
		Name:  req.Name,
		Owner: req.Owner,
	}, nil
}

func StartGRPCServer(port string) {
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	server := grpc.NewServer()
	pb.RegisterCompanyServiceServer(server, NewCompanyServiceServer())

	log.Printf("gRPC server is listening on port %s", port)
	if err := server.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
