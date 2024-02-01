package database

import "grpc-template/internal/models"

func AutoMigrate(connection IConnection) error {
	err := connection.GetMaster().AutoMigrate(
		&models.User{},
	)

	return err
}
