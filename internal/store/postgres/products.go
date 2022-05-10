package postgres

import (
	"context"

	"github.com/MTursynbekov/ElectronicsStore_gRPC/internal/models"
	"github.com/MTursynbekov/ElectronicsStore_gRPC/internal/store"
	"github.com/jmoiron/sqlx"
)

func (db *DB) Products() store.ProductsRepository {
	if db.products == nil {
		db.products = NewProductsRepository(db.conn)
	}

	return db.products
}

type ProductsRepository struct {
	conn *sqlx.DB
}

func NewProductsRepository(conn *sqlx.DB) store.ProductsRepository {
	return &ProductsRepository{conn: conn}
}

func (r ProductsRepository) Create(ctx context.Context, product *models.Product) error {
	_, err := r.conn.Exec("INSERT INTO products(name, price, category_id, brand_id, description) VALUES ($1, $2, $3, $4, $5)",
		product.Name,
		product.Price,
		product.CategoryId,
		product.BrandId,
		product.Description,
	)
	if err != nil {
		return err
	}

	return nil
}

func (r ProductsRepository) All(ctx context.Context) ([]*models.Product, error) {
	products := make([]*models.Product, 0)
	query := "SELECT * FROM products"

	if err := r.conn.Select(&products, query); err != nil {
		return nil, err
	}

	return products, nil
}

func (r ProductsRepository) ByID(ctx context.Context, id uint) (*models.Product, error) {
	product := new(models.Product)
	if err := r.conn.Get(product, "SELECT * FROM products WHERE id = $1", id); err != nil {
		return nil, err
	}

	return product, nil
}

func (r ProductsRepository) Update(ctx context.Context, product *models.Product, id uint) error {
	_, err := r.conn.
		Exec("UPDATE products SET name = $1, price = $2, category_id = $3, brand_id = $4, description = $5 WHERE id = $6",
			product.Name,
			product.Price,
			product.CategoryId,
			product.BrandId,
			product.Description,
			id,
		)
	if err != nil {
		return err
	}

	return nil
}

func (r ProductsRepository) Delete(ctx context.Context, id uint) error {
	_, err := r.conn.Exec("DELETE FROM products WHERE id = $1", id)
	if err != nil {
		return err
	}

	return nil
}
