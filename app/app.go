package app

import (
	"fmt"
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
}

func New() *App {
	return &App{
		server: grpc.NewServer(),
	}
}

func (a *App) Run(port string, timeout time.Duration) error {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		return fmt.Errorf("failed to listen: %v", err)
	}

	// Start the server in a separate goroutine
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
