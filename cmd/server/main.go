package main

import (
	"context"
	"github.com/MTursynbekov/ElectronicsStore_gRPC/internal/http"
	"github.com/MTursynbekov/ElectronicsStore_gRPC/internal/store"
	"github.com/MTursynbekov/ElectronicsStore_gRPC/internal/store/postgres"
	"log"
	"os"
	"os/signal"
	"syscall"
)

const (
	dbUrl    = "postgres://postgres:5500dake@localhost:5432/store"
	grpcPort = ":8080"
	restPort = ":8081"
)

type app struct {
	restServer http.RestServer
	grpcServer http.GRPCServer
}

func (a app) start() {
	a.restServer.Start()
	a.grpcServer.Start()
}

func (a app) shutdown() error {
	a.grpcServer.Stop()
	return a.restServer.Stop()
}

func newApp(store store.Store) (app, error) {
	service := http.NewGRPCService(store)

	gs, err := http.NewGRPCServer(service, grpcPort)
	if err != nil {
		return app{}, err
	}

	return app{
		restServer: http.NewRestServer(service, restPort),
		grpcServer: gs,
	}, nil
}

func run(ctx context.Context) error {
	s := postgres.NewDB()
	if err := s.Connect(dbUrl); err != nil {
		log.Fatalf("failed to connect to database %v: %v", dbUrl, err)
		return err
	}
	defer s.Close()

	app, err := newApp(s)
	if err != nil {
		return err
	}

	app.start()
	defer app.shutdown()

	select {
	case restErr := <-app.restServer.Error():
		return restErr
	case grpcErr := <-app.grpcServer.Error():
		return grpcErr
	case <-ctx.Done():
		return nil
	}
}

func main() {

	ctx, stop := signal.NotifyContext(context.Background(), []os.Signal{os.Interrupt, syscall.SIGTERM}...)
	defer stop()

	if err := run(ctx); err != nil {
		stop()
		log.Fatal(err)
	}
}
