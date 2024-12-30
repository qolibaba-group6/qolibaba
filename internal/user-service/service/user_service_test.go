
package service_test

import (
	"context"
	"testing"
	"time"
	"travel-booking-app/internal/user-service/service"
	pb "travel-booking-app/internal/user-service"

	"github.com/stretchr/testify/assert"
)

func TestGetUser_ValidID(t *testing.T) {
	// Arrange
	server := &service.UserServiceServer{}
	req := &pb.GetUserRequest{Id: "123"}

	// Act
	resp, err := server.GetUser(context.Background(), req)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, "Advanced User", resp.Name)
}

func TestGetUser_InvalidID(t *testing.T) {
	// Arrange
	server := &service.UserServiceServer{}
	req := &pb.GetUserRequest{Id: "999"}

	// Act
	resp, err := server.GetUser(context.Background(), req)

	// Assert
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Equal(t, "user not found", err.Error())
}

func TestGetUser_MissingID(t *testing.T) {
	// Arrange
	server := &service.UserServiceServer{}
	req := &pb.GetUserRequest{Id: ""}

	// Act
	resp, err := server.GetUser(context.Background(), req)

	// Assert
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Equal(t, "user ID is required", err.Error())
}

func TestGetUser_ContextTimeout(t *testing.T) {
	// Arrange
	server := &service.UserServiceServer{}
	req := &pb.GetUserRequest{Id: "123"}
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	// Act
	resp, err := server.GetUser(ctx, req)

	// Assert
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Equal(t, context.DeadlineExceeded, err)
}
