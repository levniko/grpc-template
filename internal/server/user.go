package server

import (
	"context"
	users "grpc-template/internal/modules/user"
	"grpc-template/protobuf/user"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserController struct {
	userService users.Service
}

func NewUserController(userService users.Service) *UserController {
	return &UserController{
		userService: userService,
	}
}

// CreateUser, is an imlementation of gRPC CreateUser RPC method
func (u *UserController) CreateUser(ctx context.Context, req *user.CreateUserRequest) (*user.CreateUserResponse, error) {
	createdUser, err := u.userService.CreateUser(ctx, req.GetUser())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create user: %v", err)
	}

	response := &user.CreateUserResponse{
		User: &user.UserResponse{
			FirstName: createdUser.FirstName,
			LastName:  createdUser.LastName,
			Email:     createdUser.Email,
		},
	}

	return response, nil
}
