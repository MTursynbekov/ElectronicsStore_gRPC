package gRPC_services

import (
	"context"
	"github.com/MTursynbekov/ElectronicsStore_gRPC/api"
	"github.com/MTursynbekov/ElectronicsStore_gRPC/internal/models"
	"github.com/MTursynbekov/ElectronicsStore_gRPC/internal/store"
)

type CategoriesService struct {
	Store store.Store

	api.UnimplementedCategoryServiceServer
}

func (s *CategoriesService) GetCategoryList(ctx context.Context, empty *api.Empty) (*api.CategoriesResponse, error) {
	categoriesResponse := new(api.CategoriesResponse)
	categories, err := s.Store.Categories().All(ctx)
	if err != nil {
		return nil, err
	}

	for _, category := range categories {
		categoriesResponse.Categories = append(categoriesResponse.Categories, &api.Category{
			Id:   int64(category.ID),
			Name: category.Name,
		})
	}

	return categoriesResponse, nil
}

func (s *CategoriesService) GetCategoryById(ctx context.Context, id *api.IdRequest) (*api.Category, error) {
	category, err := s.Store.Categories().ByID(ctx, uint(id.Id))
	if err != nil {
		return nil, err
	}

	categoryResponse := &api.Category{
		Id:   int64(category.ID),
		Name: category.Name,
	}

	return categoryResponse, nil
}

func (s *CategoriesService) CreateCategory(ctx context.Context, categoryRequest *api.CategoryRequest) (*api.Empty, error) {
	category := &models.Category{
		Name: categoryRequest.Name,
	}

	err := s.Store.Categories().Create(ctx, category)
	if err != nil {
		return nil, err
	}

	return &api.Empty{}, nil
}

func (s *CategoriesService) UpdateCategory(ctx context.Context, categoryRequest *api.Category) (*api.Empty, error) {
	category := &models.Category{
		Name: categoryRequest.Name,
	}

	err := s.Store.Categories().Update(ctx, category, uint(categoryRequest.Id))
	if err != nil {
		return nil, err
	}

	return &api.Empty{}, nil
}

func (s *CategoriesService) DeleteCategory(ctx context.Context, id *api.IdRequest) (*api.Empty, error) {
	err := s.Store.Categories().Delete(ctx, uint(id.Id))
	if err != nil {
		return nil, err
	}

	return &api.Empty{}, nil
}
