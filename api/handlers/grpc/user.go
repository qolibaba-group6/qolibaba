package grpc

import (
	"context"
	"log"

	pb "github.com/ehsansobhani/project_structure-3/api/pb"
	"github.com/ehsansobhani/project_structure-3/app/user"
	"github.com/ehsansobhani/project_structure-3/internal/user/domain"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserService struct {
	pb.UnimplementedUserServiceServer
	AppService user.UserAppService
}

func NewUserService(s user.UserAppService) *UserService {
	return &UserService{AppService: s}
}

func (s *UserService) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	user := &domain.User{
		Name:     req.GetName(),
		Email:    req.GetEmail(),
		Password: req.GetPassword(),
	}

	err := s.AppService.Register(user)
	if err != nil {
		log.Printf("Register error: %v", err)
		return nil, status.Errorf(codes.Internal, "could not register user: %v", err)
	}

	return &pb.RegisterResponse{Id: user.ID.String()}, nil
}

func (s *UserService) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	token, err := s.AppService.Login(req.GetEmail(), req.GetPassword())
	if err != nil {
		log.Printf("Login error: %v", err)
		return nil, status.Errorf(codes.Unauthenticated, "could not login: %v", err)
	}

	return &pb.LoginResponse{
		Token: token,
	}, nil
}

func (s *UserService) GetUserProfile(ctx context.Context, req *pb.GetUserProfileRequest) (*pb.GetUserProfileResponse, error) {
	userID, err := uuid.Parse(req.GetId())
	if err != nil {
		log.Printf("Invalid UUID: %v", err)
		return nil, status.Errorf(codes.InvalidArgument, "invalid user ID: %v", err)
	}

	user, err := s.AppService.GetProfile(userID)
	if err != nil {
		log.Printf("GetUserProfile error: %v", err)
		return nil, status.Errorf(codes.NotFound, "could not find user: %v", err)
	}

	return &pb.GetUserProfileResponse{
		Id:        user.ID.String(),
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt.Unix(),
		UpdatedAt: user.UpdatedAt.Unix(),
	}, nil
}

func (s *UserService) UpdateUserProfile(ctx context.Context, req *pb.UpdateUserProfileRequest) (*pb.UpdateUserProfileResponse, error) {
	userID, err := uuid.Parse(req.GetId())
	if err != nil {
		log.Printf("Invalid UUID: %v", err)
		return nil, status.Errorf(codes.InvalidArgument, "invalid user ID: %v", err)
	}

	user, err := s.AppService.GetProfile(userID)
	if err != nil {
		log.Printf("GetUserProfile error: %v", err)
		return nil, status.Errorf(codes.NotFound, "could not find user: %v", err)
	}

	if req.GetName() != "" {
		user.Name = req.GetName()
	}
	if req.GetEmail() != "" {
		user.Email = req.GetEmail()
	}
	if req.GetPassword() != "" {
		user.Password = req.GetPassword()
	}

	err = s.AppService.UpdateProfile(user)
	if err != nil {
		log.Printf("UpdateUserProfile error: %v", err)
		return nil, status.Errorf(codes.Internal, "could not update user: %v", err)
	}

	return &pb.UpdateUserProfileResponse{
		Id:        user.ID.String(),
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt.Unix(),
		UpdatedAt: user.UpdatedAt.Unix(),
	}, nil
}

func (s *UserService) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	userID, err := uuid.Parse(req.GetId())
	if err != nil {
		log.Printf("Invalid UUID: %v", err)
		return nil, status.Errorf(codes.InvalidArgument, "invalid user ID: %v", err)
	}

	err = s.AppService.DeleteUser(userID)
	if err != nil {
		log.Printf("DeleteUser error: %v", err)
		return nil, status.Errorf(codes.Internal, "could not delete user: %v", err)
	}

	return &pb.DeleteUserResponse{Success: true}, nil
}
