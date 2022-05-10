package store

import (
	"context"

	"github.com/MTursynbekov/ElectronicsStore_gRPC/internal/models"
)

type Store interface {
	Connect(url string) error
	Close() error

	Categories() CategoriesRepository
	Brands() BrandsRepository
	Products() ProductsRepository
	ProductSpecifications() ProductSpecificationsRepository
	ProductImages() ProductImagesRepository
}

type CategoriesRepository interface {
	Create(ctx context.Context, category *models.Category) error
	All(ctx context.Context) ([]*models.Category, error)
	ByID(ctx context.Context, id uint) (*models.Category, error)
	Update(ctx context.Context, category *models.Category, id uint) error
	Delete(ctx context.Context, id uint) error
}

type BrandsRepository interface {
	Create(ctx context.Context, brand *models.Brand) error
	All(ctx context.Context) ([]*models.Brand, error)
	ByID(ctx context.Context, id uint) (*models.Brand, error)
	Update(ctx context.Context, brand *models.Brand, id uint) error
	Delete(ctx context.Context, id uint) error
}

type ProductsRepository interface {
	Create(ctx context.Context, product *models.Product) error
	All(ctx context.Context) ([]*models.Product, error)
	ByID(ctx context.Context, id uint) (*models.Product, error)
	Update(ctx context.Context, product *models.Product, id uint) error
	Delete(ctx context.Context, id uint) error
}

type ProductSpecificationsRepository interface {
	Create(ctx context.Context, productSpecification *models.ProductSpecification) error
	All(ctx context.Context, productId uint) ([]*models.ProductSpecification, error)
	Delete(ctx context.Context, id uint) error
}

type ProductImagesRepository interface {
	Create(ctx context.Context, productImage *models.ProductImage) error
	All(ctx context.Context, productId uint) ([]*models.ProductImage, error)
	Delete(ctx context.Context, id uint) error
}
