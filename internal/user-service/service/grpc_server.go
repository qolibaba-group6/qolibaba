package service

import (
	"context"
	"errors"
	"log"
	"net"
	"time"

	// Import the generated pb package correctly
	pb "travel-booking-app/internal/user-service" // Adjust the path if necessary

	"google.golang.org/grpc"
)

// UserServiceServer implements the UserServiceServer interface defined in the proto file
type UserServiceServer struct {
	pb.UnimplementedUserServiceServer
}

// GetUser retrieves user information by ID with error handling and validation
func (s *UserServiceServer) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	// Validate input
	if req.Id == "" {
		log.Println("Invalid request: missing user ID")
		return nil, errors.New("user ID is required")
	}

	// Simulated database fetch with delay to demonstrate context timeout
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

// StartGRPCServer initializes and starts the gRPC server with logging middleware
func StartGRPCServer() {
	// Listen on TCP port 50051
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// Create a new gRPC server instance with a logging interceptor
	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(loggingInterceptor), // Adding middleware for logging
	)

	// Register the UserServiceServer implementation
	pb.RegisterUserServiceServer(grpcServer, &UserServiceServer{})
	log.Println("gRPC server is running on port 50051")

	// Serve gRPC requests
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

// loggingInterceptor is a middleware that logs requests and responses
func loggingInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	// Log the gRPC method and request
	log.Printf("gRPC method: %s, request: %v", info.FullMethod, req)
	resp, err := handler(ctx, req)
	if err != nil {
		// Log the error if any
		log.Printf("Error: %v", err)
	} else {
		// Log the response if no error occurred
		log.Printf("Response: %v", resp)
	}
	return resp, err
}
