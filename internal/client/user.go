package client

import (
	"context"
	"fmt"
	"grpc-template/protobuf/user"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type CreateUserInput struct {
	FirstName     string
	LastName      string
	Email         string
	Password      string
	PasswordAgain string
}

type UserClient struct {
	client user.UserServiceClient
}

func NewUserClient(address string) *UserClient {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock(), grpc.WithTimeout(5*time.Second))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	return &UserClient{
		client: user.NewUserServiceClient(conn),
	}
}

// CreateUser, sunucu üzerinde bir kullanıcı oluşturmak için bir istemci fonksiyonudur.
func (c *UserClient) CreateUser(ctx context.Context, input CreateUserInput) {
	req := &user.CreateUserRequest{
		User: &user.UserRequest{
			FirstName:     input.FirstName,
			LastName:      input.LastName,
			Email:         input.Email,
			Password:      input.Password,
			PasswordAgain: input.PasswordAgain,
			//Address:       input.Address,
		},
	}
	resp, err := c.client.CreateUser(ctx, req)
	if err != nil {
		log.Fatalf("could not create user: %v", err)
	}
	fmt.Printf("Created user: %v\n", resp.GetUser())
}
