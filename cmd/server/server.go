package main

import (
	"log"
	"net"

	"github.com/MTursynbekov/ElectronicsStore_gRPC/api"

	"github.com/MTursynbekov/ElectronicsStore_gRPC/internal/store/postgres"

	"google.golang.org/grpc"
)

const (
	dbUrl = "postgres://postgres:5500dake@localhost:5432/store"
	port  = ":8080"
)

func main() {
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("cannot listen to %s: %v", port, err)
	}
	defer listener.Close()

	store := postgres.NewDB()
	if err = store.Connect(dbUrl); err != nil {
		log.Fatalf("failed to connect to database: %v", dbUrl)
	}
	defer store.Close()

	grpcServer := grpc.NewServer()

	categoriesService := &CategoriesService{
		store: store,
	}
	brandsService := &BrandsService{
		store: store,
	}
	productsService := &ProductsService{
		store: store,
	}

	api.RegisterCategoryServiceServer(grpcServer, categoriesService)
	api.RegisterBrandServiceServer(grpcServer, brandsService)
	api.RegisterProductServiceServer(grpcServer, productsService)

	log.Printf("Serving on %v", listener.Addr())
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve on %v: %v", listener.Addr(), err)
	}
}
