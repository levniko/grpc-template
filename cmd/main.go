package main

import (
	"grpc-template/app"
	"log"
	"time"
)

func main() {
	app, err := app.New()
	if err != nil {
		log.Fatalf("Error creating new app: %v", err)
	}
	if err := app.Run(":50051", 5*time.Second); err != nil {
		log.Fatalf("Error running server: %v", err)
	}
}
