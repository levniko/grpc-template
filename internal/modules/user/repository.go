package users

import (
	"context"
	"grpc-template/internal/database"
	"grpc-template/internal/models"
)

type IRepository interface {
	Save(ctx context.Context, user *models.User) error
	FindByEmail(ctx context.Context, email string) (*models.User, error)
	FindByPhoneNumber(ctx context.Context, phoneNumber string) (*models.User, error)
}

type Repository struct {
	Connection database.IConnection
}

func NewUserRepository(c database.IConnection) *Repository {
	return &Repository{Connection: c}
}

func (r *Repository) Save(ctx context.Context, mdl *models.User) error {
	return r.Connection.GetMaster().Create(mdl).Error
}

func (r *Repository) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User
	err := r.Connection.GetSlave().Where("email = ?", email).Limit(1).Take(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *Repository) FindByPhoneNumber(ctx context.Context, phoneNumber string) (*models.User, error) {
	var user models.User
	err := r.Connection.GetSlave().Where("phone_number = ?", phoneNumber).Limit(1).Take(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
