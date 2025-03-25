package app

import (
	climbgrpc "github.com/didsqq/crud-service-alpinizm/internal/grpc"

	"google.golang.org/grpc"
)

type App struct {
	gRPCServer *grpc.Server
	port       int
}

func New(
	service climbgrpc.Climbs,
	port int,
) *App {
	gRPCServer := grpc.NewServer()

	climbgrpc.Register(gRPCServer, service)

	return &App{
		gRPCServer: gRPCServer,
		port:       port,
	}
}
