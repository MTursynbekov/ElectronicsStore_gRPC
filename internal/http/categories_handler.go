package http

import (
	"github.com/MTursynbekov/ElectronicsStore_gRPC/api"
	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/jsonpb"
	"net/http"
	"strconv"
)

func (s RestServer) GetCategoryListHandler(c *gin.Context) {
	req := new(api.Empty)
	resp, err := s.service.categoriesService.GetCategoryList(c, req)
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

func (s RestServer) CreateCategoryHandler(c *gin.Context) {
	var req api.CategoryRequest

	err := jsonpb.Unmarshal(c.Request.Body, &req)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	resp, err := s.service.categoriesService.CreateCategory(c, &req)
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

func (s RestServer) GetCategoryHandler(c *gin.Context) {
	p := c.Param("id")
	if p == "" {
		c.JSON(http.StatusBadRequest, "id param can not be empty")
		return
	}

	id, _ := strconv.ParseInt(p, 10, 64)
	req := &api.IdRequest{
		Id: id,
	}

	resp, err := s.service.categoriesService.GetCategoryById(c, req)
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

func (s RestServer) UpdateCategoryHandler(c *gin.Context) {
	p := c.Param("id")
	if p == "" {
		c.JSON(http.StatusBadRequest, "id param can not be empty")
		return
	}

	id, _ := strconv.ParseInt(p, 10, 64)

	req := &api.Category{
		Id: id,
	}
	err := jsonpb.Unmarshal(c.Request.Body, req)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	resp, err := s.service.categoriesService.UpdateCategory(c, req)
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

func (s RestServer) DeleteCategoryHandler(c *gin.Context) {
	p := c.Param("id")
	if p == "" {
		c.JSON(http.StatusBadRequest, "id param can not be empty")
		return
	}

	id, _ := strconv.ParseInt(p, 10, 64)
	req := &api.IdRequest{
		Id: id,
	}

	resp, err := s.service.categoriesService.DeleteCategory(c, req)
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
