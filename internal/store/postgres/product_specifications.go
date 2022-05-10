package postgres

import (
	"context"

	"github.com/MTursynbekov/ElectronicsStore_gRPC/internal/models"
	"github.com/MTursynbekov/ElectronicsStore_gRPC/internal/store"
	"github.com/jmoiron/sqlx"
)

func (db *DB) ProductSpecifications() store.ProductSpecificationsRepository {
	if db.productSpecifications == nil {
		db.productSpecifications = NewProductSpecificationsRepository(db.conn)
	}

	return db.productSpecifications
}

type ProductSpecificationsRepository struct {
	conn *sqlx.DB
}

func NewProductSpecificationsRepository(conn *sqlx.DB) store.ProductSpecificationsRepository {
	return &ProductSpecificationsRepository{conn: conn}
}

func (r ProductSpecificationsRepository) Create(ctx context.Context, productSpecification *models.ProductSpecification) error {
	_, err := r.conn.Exec("INSERT INTO product_specifications(key, value, product_id) VALUES ($1, $2, $3)",
		productSpecification.Key,
		productSpecification.Value,
		productSpecification.ProductId,
	)
	if err != nil {
		return err
	}

	return nil
}

func (r ProductSpecificationsRepository) All(ctx context.Context, productId uint) ([]*models.ProductSpecification, error) {
	productSpecifications := make([]*models.ProductSpecification, 0)
	query := "SELECT * FROM product_specifications WHERE product_id = $1"

	if err := r.conn.Select(&productSpecifications, query, productId); err != nil {
		return nil, err
	}

	return productSpecifications, nil
}

func (r ProductSpecificationsRepository) Delete(ctx context.Context, id uint) error {
	_, err := r.conn.Exec("DELETE FROM product_specifications WHERE id = $1", id)
	if err != nil {
		return err
	}

	return nil
}
