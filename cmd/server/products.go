package main

import (
	"context"

	"github.com/MTursynbekov/ElectronicsStore_gRPC/api"
	"github.com/MTursynbekov/ElectronicsStore_gRPC/internal/models"
	"github.com/MTursynbekov/ElectronicsStore_gRPC/internal/store"
)

type ProductsService struct {
	store store.Store

	api.UnimplementedProductServiceServer
}

func (s *ProductsService) GetProductList(ctx context.Context, empty *api.Empty) (*api.ProductsResponse, error) {
	productsResponse := new(api.ProductsResponse)
	products, err := s.store.Products().All(ctx)
	if err != nil {
		return nil, err
	}

	for _, product := range products {
		productResponse := &api.Product{
			Id:          int64(product.ID),
			Name:        product.Name,
			Price:       product.Price,
			Description: product.Description,
		}

		err = setProductDependencies(ctx, s.store, productResponse, product)
		if err != nil {
			return nil, err
		}

		productsResponse.Products = append(productsResponse.Products, productResponse)
	}

	return productsResponse, nil
}

func (s *ProductsService) GetProductById(ctx context.Context, id *api.IdRequest) (*api.Product, error) {
	product, err := s.store.Products().ByID(ctx, uint(id.Id))
	if err != nil {
		return nil, err
	}

	productResponse := &api.Product{
		Id:          int64(product.ID),
		Name:        product.Name,
		Price:       product.Price,
		Description: product.Description,
	}

	err = setProductDependencies(ctx, s.store, productResponse, product)
	if err != nil {
		return nil, err
	}

	return productResponse, nil
}

func (s *ProductsService) CreateProduct(ctx context.Context, productResponse *api.ProductRequest) (*api.Empty, error) {
	product := &models.Product{
		Name:        productResponse.Name,
		Price:       productResponse.Price,
		CategoryId:  uint(productResponse.CategoryId),
		BrandId:     uint(productResponse.BrandId),
		Description: productResponse.Description,
	}

	err := s.store.Products().Create(ctx, product)
	if err != nil {
		return nil, err
	}

	for _, spec := range productResponse.Specifications {
		err = s.store.ProductSpecifications().Create(ctx, &models.ProductSpecification{
			Key:   spec.Key,
			Value: spec.Value,
		})
		if err != nil {
			return nil, err
		}
	}

	for _, img := range productResponse.Images {
		err = s.store.ProductImages().Create(ctx, &models.ProductImage{
			Src: img,
		})
		if err != nil {
			return nil, err
		}
	}

	return &api.Empty{}, nil
}

func (s *ProductsService) UpdateProduct(ctx context.Context, productResponse *api.ProductUpdateRequest) (*api.Empty, error) {
	product := &models.Product{
		Name:        productResponse.Product.Name,
		Price:       productResponse.Product.Price,
		CategoryId:  uint(productResponse.Product.CategoryId),
		BrandId:     uint(productResponse.Product.BrandId),
		Description: productResponse.Product.Description,
	}

	err := s.store.Products().Update(ctx, product, uint(productResponse.Id))
	if err != nil {
		return nil, err
	}

	return &api.Empty{}, nil
}

func (s *ProductsService) DeleteProduct(ctx context.Context, id *api.IdRequest) (*api.Empty, error) {
	err := s.store.Products().Delete(ctx, uint(id.Id))
	if err != nil {
		return nil, err
	}

	return &api.Empty{}, nil
}

func setProductDependencies(ctx context.Context, s store.Store, productResponse *api.Product, product *models.Product) error {
	category, err := s.Categories().ByID(ctx, uint(product.CategoryId))
	if err != nil {
		return err
	}

	brand, err := s.Brands().ByID(ctx, uint(product.BrandId))
	if err != nil {
		return err
	}

	specs, err := s.ProductSpecifications().All(ctx, uint(product.ID))
	if err != nil {
		return err
	}

	specsResponse := make([]*api.ProductSpecification, 0)
	for _, spec := range specs {
		specsResponse = append(specsResponse, &api.ProductSpecification{
			Key:   spec.Key,
			Value: spec.Value,
		})
	}

	imgs, err := s.ProductImages().All(ctx, uint(product.ID))
	if err != nil {
		return err
	}

	imgsResponse := make([]string, 0)
	for _, img := range imgs {
		imgsResponse = append(imgsResponse, img.Src)
	}

	productResponse.Category = &api.Category{
		Id:   int64(category.ID),
		Name: category.Name,
	}

	productResponse.Brand = &api.Brand{
		Id:   int64(brand.ID),
		Name: brand.Name,
	}

	productResponse.Specifications = specsResponse
	productResponse.Images = imgsResponse

	return nil
}
