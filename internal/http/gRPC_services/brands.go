package gRPC_services

import (
	"context"
	"github.com/MTursynbekov/ElectronicsStore_gRPC/api"
	"github.com/MTursynbekov/ElectronicsStore_gRPC/internal/models"
	"github.com/MTursynbekov/ElectronicsStore_gRPC/internal/store"
)

type BrandsService struct {
	Store store.Store

	api.UnimplementedBrandServiceServer
}

func (s *BrandsService) GetBrandList(ctx context.Context, empty *api.Empty) (*api.BrandsResponse, error) {
	brandsResponse := new(api.BrandsResponse)
	brands, err := s.Store.Brands().All(ctx)
	if err != nil {
		return nil, err
	}

	for _, brand := range brands {
		brandsResponse.Brands = append(brandsResponse.Brands, &api.Brand{
			Id:   int64(brand.ID),
			Name: brand.Name,
		})
	}

	return brandsResponse, nil
}

func (s *BrandsService) GetBrandById(ctx context.Context, id *api.IdRequest) (*api.Brand, error) {
	brand, err := s.Store.Brands().ByID(ctx, uint(id.Id))
	if err != nil {
		return nil, err
	}

	brandResponse := &api.Brand{
		Id:   int64(brand.ID),
		Name: brand.Name,
	}

	return brandResponse, nil
}

func (s *BrandsService) CreateBrand(ctx context.Context, brandRequest *api.BrandRequest) (*api.Empty, error) {
	brand := &models.Brand{
		Name: brandRequest.Name,
	}

	err := s.Store.Brands().Create(ctx, brand)
	if err != nil {
		return nil, err
	}

	return &api.Empty{}, nil
}

func (s *BrandsService) UpdateBrand(ctx context.Context, brandRequest *api.Brand) (*api.Empty, error) {
	brand := &models.Brand{
		Name: brandRequest.Name,
	}

	err := s.Store.Brands().Update(ctx, brand, uint(brandRequest.Id))
	if err != nil {
		return nil, err
	}

	return &api.Empty{}, nil
}

func (s *BrandsService) DeleteBrand(ctx context.Context, id *api.IdRequest) (*api.Empty, error) {
	err := s.Store.Brands().Delete(ctx, uint(id.Id))
	if err != nil {
		return nil, err
	}

	return &api.Empty{}, nil
}
