package http

import (
	"github.com/MTursynbekov/ElectronicsStore_gRPC/api"
	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/jsonpb"
	"net/http"
	"strconv"
)

func (s RestServer) GetProductListHandler(c *gin.Context) {
	req := new(api.Empty)
	resp, err := s.service.productsService.GetProductList(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	m := new(jsonpb.Marshaler)
	if err = m.Marshal(c.Writer, resp); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
}

func (s RestServer) CreateProductHandler(c *gin.Context) {
	var req api.ProductRequest

	err := jsonpb.Unmarshal(c.Request.Body, &req)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	resp, err := s.service.productsService.CreateProduct(c, &req)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	m := &jsonpb.Marshaler{}
	if err := m.Marshal(c.Writer, resp); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
}

func (s RestServer) GetProductHandler(c *gin.Context) {
	p := c.Param("id")
	if p == "" {
		c.JSON(http.StatusBadRequest, "id param can not be empty")
		return
	}

	id, _ := strconv.ParseInt(p, 10, 64)
	req := &api.IdRequest{
		Id: id,
	}

	resp, err := s.service.productsService.GetProductById(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	m := new(jsonpb.Marshaler)
	if err = m.Marshal(c.Writer, resp); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
}

func (s RestServer) UpdateProductHandler(c *gin.Context) {
	p := c.Param("id")
	if p == "" {
		c.JSON(http.StatusBadRequest, "id param can not be empty")
		return
	}

	id, _ := strconv.ParseInt(p, 10, 64)

	req := &api.ProductUpdateRequest{
		Id: id,
	}
	err := jsonpb.Unmarshal(c.Request.Body, req)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	resp, err := s.service.productsService.UpdateProduct(c, req)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	m := &jsonpb.Marshaler{}
	if err := m.Marshal(c.Writer, resp); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
}

func (s RestServer) DeleteProductHandler(c *gin.Context) {
	p := c.Param("id")
	if p == "" {
		c.JSON(http.StatusBadRequest, "id param can not be empty")
		return
	}

	id, _ := strconv.ParseInt(p, 10, 64)
	req := &api.IdRequest{
		Id: id,
	}

	resp, err := s.service.productsService.DeleteProduct(c, req)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	m := &jsonpb.Marshaler{}
	if err := m.Marshal(c.Writer, resp); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
}
