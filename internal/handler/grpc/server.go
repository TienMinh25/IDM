package grpc

import (
	"context"
	"net"

	"github.com/TienMinh25/IDM/internal/generated/grpc/go_load"
	"github.com/TienMinh25/IDM/internal/handler/grpc"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/validator"
)

type Server interface {
	Start(ctx context.Context) error
}

type server struct {
	handler go_load.GoLoadServer
}

func NewServer(handler go_load.GoLoadServer) Server {
	return &server{
		handler: handler,
	}
}

func (s *server) Start(ctx context.Context) error {
	listener, err := net.Listen("tcp", "localhost:8080")
	
	if err != nil {
		return err
	}

	defer listener.Close()

	server := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			validator.UnaryServerInterceptor(),
		),
		grpc.ChainStreamInterceptor(
			validator.StreamServerInterceptor(),
		),
	)
	go_load.RegisterGoLoadServer(server, s.handler)
	return server.Serve(listener)
}