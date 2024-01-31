package users

import (
	"context"
	"grpc-template/internal/models"
)

type IRepository interface {
	Save(ctx context.Context, user *models.User) error
	FindByEmail(ctx context.Context, email string) (*models.User, error)
	FindByPhoneNumber(ctx context.Context, phoneNumber string) (*models.User, error)
}

type Repository struct {
	//Connection database.IConnection
}

func GetUserRepository( /*c database.IConnection*/ ) *Repository {
	return &Repository{ /*Connection: c*/ }
}

func (r *Repository) Save(ctx context.Context, mdl *models.User) error {
	return nil
}

func (r *Repository) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	return nil, nil
}

func (r *Repository) FindByPhoneNumber(ctx context.Context, phoneNumber string) (*models.User, error) {
	return nil, nil
}
