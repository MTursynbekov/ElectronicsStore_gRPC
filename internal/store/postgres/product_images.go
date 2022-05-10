package postgres

import (
	"context"
	"github.com/MTursynbekov/ElectronicsStore_gRPC/internal/models"
	"github.com/MTursynbekov/ElectronicsStore_gRPC/internal/store"
	"github.com/jmoiron/sqlx"
)

func (db *DB) ProductImages() store.ProductImagesRepository {

	if db.productImages == nil {
		db.productImages = NewProductImagesRepository(db.conn)
	}

	return db.productImages
}

type ProductImagesRepository struct {
	conn *sqlx.DB
}

func NewProductImagesRepository(conn *sqlx.DB) store.ProductImagesRepository {
	return &ProductImagesRepository{conn: conn}
}

func (r ProductImagesRepository) Create(ctx context.Context, productImage *models.ProductImage) error {
	_, err := r.conn.Exec("INSERT INTO product_images(src, product_id) VALUES ($1, $2)",
		productImage.Src,
		productImage.ProductId,
	)
	if err != nil {
		return err
	}

	return nil
}

func (r ProductImagesRepository) All(ctx context.Context, productId uint) ([]*models.ProductImage, error) {
	productImages := make([]*models.ProductImage, 0)
	query := "SELECT * FROM product_images WHERE product_id = $1"

	if err := r.conn.Select(&productImages, query, productId); err != nil {
		return nil, err
	}

	return productImages, nil
}

func (r ProductImagesRepository) Delete(ctx context.Context, id uint) error {
	_, err := r.conn.Exec("DELETE FROM product_images WHERE id = $1", id)
	if err != nil {
		return err
	}

	return nil
}
