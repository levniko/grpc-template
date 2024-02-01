package app

import (
	"fmt"
	"grpc-template/internal/config"
	"grpc-template/internal/database"
	users "grpc-template/internal/modules/user"
	"grpc-template/internal/server"
	"grpc-template/protobuf/user"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"google.golang.org/grpc"
)

type App struct {
	server *grpc.Server
	config config.Config
}

func New() (*App, error) {
	cfg, err := config.New("./config.toml")
	if err != nil {
		return nil, err
	}
	return &App{
		server: grpc.NewServer(),
		config: *cfg,
	}, nil
}

func (a *App) Run(port string, timeout time.Duration) error {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		return fmt.Errorf("failed to listen: %v", err)
	}
	db, _ := database.NewDBManager(&a.config)
	userRepo := users.NewUserRepository(db)
	userUsecase := users.NewUserUsecase(userRepo)
	userServiceServer := server.NewUserServiceServer(*userUsecase)
	user.RegisterUserServiceServer(a.server, userServiceServer)

	go func() {
		if err := a.server.Serve(lis); err != nil {
			log.Printf("failed to serve: %v", err)
		}
	}()

	// Wait for interrupt signal
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	<-c

	log.Println("gRPC Server is shutting down...")

	// Initiate graceful shutdown
	a.server.GracefulStop()

	return nil
}
