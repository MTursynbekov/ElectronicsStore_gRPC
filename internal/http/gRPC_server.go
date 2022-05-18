package http

import (
	"fmt"
	"github.com/MTursynbekov/ElectronicsStore_gRPC/api"
	"github.com/MTursynbekov/ElectronicsStore_gRPC/internal/http/gRPC_services"
	"github.com/MTursynbekov/ElectronicsStore_gRPC/internal/store"
	"google.golang.org/grpc"
	"log"
	"net"
)

type GRPCService struct {
	categoriesService *gRPC_services.CategoriesService
	brandsService     *gRPC_services.BrandsService
	productsService   *gRPC_services.ProductsService
}

func NewGRPCService(store store.Store) *GRPCService {
	return &GRPCService{
		categoriesService: &gRPC_services.CategoriesService{
			Store: store,
		},
		brandsService: &gRPC_services.BrandsService{
			Store: store,
		},
		productsService: &gRPC_services.ProductsService{
			Store: store,
		},
	}
}

type GRPCServer struct {
	server   *grpc.Server
	errCh    chan error
	listener net.Listener
}

func NewGRPCServer(service *GRPCService, port string) (GRPCServer, error) {
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("cannot listen to %s: %v", port, err)
		return GRPCServer{}, err
	}
	defer listener.Close()

	server := grpc.NewServer()
	api.RegisterCategoryServiceServer(server, service.categoriesService)
	api.RegisterBrandServiceServer(server, service.brandsService)
	api.RegisterProductServiceServer(server, service.productsService)

	return GRPCServer{
		server:   server,
		listener: listener,
		errCh:    make(chan error, 1),
	}, nil
}

func (s GRPCServer) Start() {
	go func() {
		err := s.server.Serve(s.listener)
		if err != nil {
			if err.Error() == "accept tcp [::]:8080: use of closed network connection" {
				return
			}
			s.errCh <- fmt.Errorf("failed to serve on: %v: %v", s.listener.Addr(), err)
		}
	}()
}

func (s GRPCServer) Stop() {
	s.server.Stop()
}

func (s GRPCServer) Error() chan error {
	return s.errCh
}
