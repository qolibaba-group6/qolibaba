package service

import (
	"context"
	"qolibaba/api/pb"
	"qolibaba/internal/user"
	"qolibaba/internal/user/domain"
	userPort "qolibaba/internal/user/port"
	"qolibaba/pkg/jwt"
	timeutils "qolibaba/pkg/time"

	jwt2 "github.com/golang-jwt/jwt/v5"
)

var (
	ErrUserCreationValidation = user.ErrUserCreationValidation
	ErrUserOnCreate           = user.ErrUserOnCreate
)

type UserService struct {
	svc           userPort.Service
	authSecret    string
	expMin        uint
	refreshExpMin uint
}

func NewUserService(svc userPort.Service, authSecret string, expMin, refreshExpMin uint) *UserService {
	return &UserService{
		svc:           svc,
		authSecret:    authSecret,
		expMin:        expMin,
		refreshExpMin: refreshExpMin,
	}
}

func (s *UserService) SignUp(ctx context.Context, req *pb.UserSignUpRequest) (*pb.UserSignUpResponse, error) {
	userID, err := s.svc.Create(ctx, domain.User{
		FirstName: req.GetFirstName(),
		LastName:  req.GetLastName(),
		Email:     domain.Email(req.GetEmail()),
		Password:  req.GetPassword(),
	})

	if err != nil {
		return nil, err
	}

	access, refresh, err := s.createTokens(userID)
	if err != nil {
		return nil, err
	}

	return &pb.UserSignUpResponse{
		AccessToken:  access,
		RefreshToken: refresh,
	}, nil
}

func (s *UserService) createTokens(userID domain.UserUUID) (access, refresh string, err error) {
	access, err = jwt.CreateToken([]byte(s.authSecret), &jwt.UserClaims{
		RegisteredClaims: jwt2.RegisteredClaims{
			ExpiresAt: jwt2.NewNumericDate(timeutils.AddMinutes(s.expMin, true)),
		},
		UserID: userID,
	})
	if err != nil {
		return
	}

	refresh, err = jwt.CreateToken([]byte(s.authSecret), &jwt.UserClaims{
		RegisteredClaims: jwt2.RegisteredClaims{
			ExpiresAt: jwt2.NewNumericDate(timeutils.AddMinutes(s.refreshExpMin, true)),
		},
		UserID: userID,
	})

	if err != nil {
		return
	}

	return
}
