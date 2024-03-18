package main

import (
	"context"
	"log"
	"net"
	"net/http"

	"github.com/felixge/httpsnoop"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	v1 "github.com/marcosdias/tutorial-grpc-rest-api-go/gen/product/v1"
	"github.com/marcosdias/tutorial-grpc-rest-api-go/server"
	"github.com/soheilhy/cmux"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

func main() {
	grpcServer := grpc.NewServer()
	srv := server.New()
	v1.RegisterProductServiceServer(grpcServer, srv)
	reflection.Register(grpcServer)

	ctx := context.Background()
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := v1.RegisterProductServiceHandlerFromEndpoint(ctx, mux, "localhost:8081", opts)
	if err != nil {
		log.Fatal(err)
	}

	server := http.Server{
		Handler: withLogger(mux),
	}

	listener, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatal(err)
	}
	multiplexer := cmux.New(listener)
	httpL := multiplexer.Match(cmux.HTTP1Fast())
	grpcL := multiplexer.Match(cmux.HTTP2())

	go server.Serve(httpL)
	go grpcServer.Serve(grpcL)
	multiplexer.Serve()
}
func withLogger(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		m := httpsnoop.CaptureMetrics(handler, writer, request)
		log.Printf("http[%d]-- %s -- %s\n", m.Code, m.Duration, request.URL.Path)
	})
}
