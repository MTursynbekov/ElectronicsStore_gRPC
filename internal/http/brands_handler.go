package http

import (
	"github.com/MTursynbekov/ElectronicsStore_gRPC/api"
	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/jsonpb"
	"net/http"
	"strconv"
)

func (s RestServer) GetBrandListHandler(c *gin.Context) {
	req := new(api.Empty)
	resp, err := s.service.brandsService.GetBrandList(c, req)
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

func (s RestServer) CreateBrandHandler(c *gin.Context) {
	var req api.BrandRequest

	err := jsonpb.Unmarshal(c.Request.Body, &req)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	resp, err := s.service.brandsService.CreateBrand(c, &req)
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

func (s RestServer) GetBrandHandler(c *gin.Context) {
	p := c.Param("id")
	if p == "" {
		c.JSON(http.StatusBadRequest, "id param can not be empty")
		return
	}

	id, _ := strconv.ParseInt(p, 10, 64)
	req := &api.IdRequest{
		Id: id,
	}

	resp, err := s.service.brandsService.GetBrandById(c, req)
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

func (s RestServer) UpdateBrandHandler(c *gin.Context) {
	p := c.Param("id")
	if p == "" {
		c.JSON(http.StatusBadRequest, "id param can not be empty")
		return
	}

	id, _ := strconv.ParseInt(p, 10, 64)

	req := &api.Brand{
		Id: id,
	}
	err := jsonpb.Unmarshal(c.Request.Body, req)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	resp, err := s.service.brandsService.UpdateBrand(c, req)
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

func (s RestServer) DeleteBrandHandler(c *gin.Context) {
	p := c.Param("id")
	if p == "" {
		c.JSON(http.StatusBadRequest, "id param can not be empty")
		return
	}

	id, _ := strconv.ParseInt(p, 10, 64)
	req := &api.IdRequest{
		Id: id,
	}

	resp, err := s.service.brandsService.DeleteBrand(c, req)
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
