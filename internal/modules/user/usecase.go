package users

import (
	"context"
	"errors"
	"fmt"
	"grpc-template/internal/models"
	"grpc-template/pkg/argon2"
	"grpc-template/pkg/utils"
	"grpc-template/protobuf/user"
	"strings"
)

type Usecase struct {
	Repository IRepository
}

func NewUserUsecase(r IRepository) *Usecase {
	return &Usecase{Repository: r}
}

func (s *Usecase) CreateUser(ctx context.Context, req *user.UserRequest) (*models.User, error) {
	_, err := s.Repository.FindByEmail(ctx, strings.ToLower(req.Email))
	if err == nil {
		return nil, errors.New(utils.EmailAlreadyExist)
	}

	_, err = s.Repository.FindByPhoneNumber(ctx, req.PhoneNumber)
	if err == nil {
		return nil, errors.New(utils.PhoneNumberAlreadyExist)
	}

	password_config := argon2.PasswordConfig{
		Time:    1,
		Memory:  64 * 1024,
		Threads: 4,
		KeyLen:  32,
	}

	hashedPassword, err := argon2.GeneratePassword(&password_config, req.Password)
	if err != nil {
		return nil, errors.New(utils.PasswordCanNotGenerated)
	}

	model := &models.User{
		FirstName:    req.FirstName,
		LastName:     req.LastName,
		PhoneNumber:  req.PhoneNumber,
		Email:        req.Email,
		IsAdmin:      false,
		Password:     hashedPassword,
		PasswordHash: hashedPassword,
	}

	err = s.Repository.Save(ctx, model)
	if err != nil {
		return nil, fmt.Errorf("failed to create user in the repository: %v", err)
	}

	return model, nil
}
