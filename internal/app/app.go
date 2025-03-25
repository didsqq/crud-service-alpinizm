package app

import (
	climbgrpc "crud-service-alpinizm/internal/grpc"

	"google.golang.org/grpc"
)

type App struct {
	gRPCServer *grpc.Server
	port       int
}

func New(
	service climbgrpc.Server,
	port int,
) *App {
	gRPCServer := grpc.NewServer()

	climbgrpc.Register(gRPCServer, service)

	return &App{
		gRPCServer: gRPCServer,
		port:       port,
	}
}
