package main

import (
	"fmt"
	"log"
	"net"
	"os"

	v1 "github.com/marcosdias/tutorial-grpc-rest-api-go/gen/product/v1"
	"github.com/marcosdias/tutorial-grpc-rest-api-go/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	grpcServer := grpc.NewServer()
	srv := server.New()
	v1.RegisterProductServiceServer(grpcServer, srv)
	reflection.Register(grpcServer)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return fmt.Errorf("failed to listen to port %s", port)
	}
	defer listener.Close()

	return grpcServer.Serve(listener)
}
