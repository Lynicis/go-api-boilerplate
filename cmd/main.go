package main

import (
	"log"
	"turkic-mythology-gateway/cmd/server"
)

func main() {
	srv := server.NewGatewayServer(":8080")
	err := srv.Start()
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
