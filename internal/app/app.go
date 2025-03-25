package app

import (
	"fmt"
	"log"
	"net"

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

func (a *App) Run() error {
	const op = "grpcapp.Run"

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	log.Println("gRPC server is running", "addr", l.Addr().String())

	if err := a.gRPCServer.Serve(l); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (a *App) Stop() {
	const op = "grpcapp.Stop"

	log.Println("stopping gRPC server")

	a.gRPCServer.GracefulStop()
}
