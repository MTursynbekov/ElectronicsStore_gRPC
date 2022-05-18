package http

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "google.golang.org/grpc"
)

type RestServer struct {
	server  *http.Server
	service *GRPCService
	errCh   chan error
}

func NewRestServer(service *GRPCService, port string) RestServer {
	router := gin.Default()

	rs := RestServer{
		server: &http.Server{
			Addr:    port,
			Handler: router,
		},
		service: service,
	}

	registerRoutes(router, rs)

	return rs
}

func registerRoutes(r *gin.Engine, rs RestServer) {
	// register categories routes
	r.GET("/categories", rs.GetCategoryListHandler)
	r.POST("/categories", rs.CreateCategoryHandler)
	r.GET("/categories/:id", rs.GetCategoryHandler)
	r.PUT("/categories/:id", rs.UpdateCategoryHandler)
	r.DELETE("/categories/:id", rs.DeleteCategoryHandler)

	// register brands routes
	r.GET("/brands", rs.GetBrandListHandler)
	r.POST("/brands", rs.CreateBrandHandler)
	r.GET("/brands/:id", rs.GetBrandHandler)
	r.PUT("/brands/:id", rs.UpdateBrandHandler)
	r.DELETE("/brands/:id", rs.DeleteBrandHandler)

	// register products routes
	r.GET("/products", rs.GetProductListHandler)
	r.POST("/products", rs.CreateProductHandler)
	r.GET("/products/:id", rs.GetProductHandler)
	r.PUT("/products/:id", rs.UpdateProductHandler)
	r.DELETE("/products/:id", rs.DeleteProductHandler)
}

func (s RestServer) Start() {
	go func() {
		if err := s.server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			s.errCh <- fmt.Errorf("failed to serve on: %v: %v", s.server.Addr, err)
		}
	}()
}

func (s RestServer) Stop() error {
	return s.server.Shutdown(context.Background())
}

func (s RestServer) Error() chan error {
	return s.errCh
}
