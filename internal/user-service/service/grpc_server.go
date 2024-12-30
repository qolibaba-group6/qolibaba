package service
import (
	"context"
	"errors"
	"log"
	"net"
	"time"
	pb "travel-booking-app/internal/user-service" 
	"google.golang.org/grpc"
)
type UserServiceServer struct {
	pb.UnimplementedUserServiceServer
}
func (s *UserServiceServer) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	if req.Id == "" {
		log.Println("Invalid request: missing user ID")
		return nil, errors.New("user ID is required")
	}
	select {
	case <-time.After(2 * time.Second):
		if req.Id == "123" {
			log.Printf("User found: %s", req.Id)
			return &pb.GetUserResponse{Name: "Advanced User"}, nil
		}
		log.Printf("User not found: %s", req.Id)
		return nil, errors.New("user not found")
	case <-ctx.Done():
		log.Println("Request cancelled or timed out")
		return nil, ctx.Err()
	}
}
func StartGRPCServer() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(loggingInterceptor), // Adding middleware for logging
	)
	pb.RegisterUserServiceServer(grpcServer, &UserServiceServer{})
	log.Println("gRPC server is running on port 50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
func loggingInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	log.Printf("gRPC method: %s, request: %v", info.FullMethod, req)
	resp, err := handler(ctx, req)
	if err != nil {
		log.Printf("Error: %v", err)
	} else {
		log.Printf("Response: %v", resp)
	}
	return resp, err
}
