package postgres

import (
	"context"

	"github.com/MTursynbekov/ElectronicsStore_gRPC/internal/models"
	"github.com/MTursynbekov/ElectronicsStore_gRPC/internal/store"
	"github.com/jmoiron/sqlx"
)

func (db *DB) Brands() store.BrandsRepository {
	if db.brands == nil {
		db.brands = NewBrandsRepository(db.conn)
	}

	return db.brands
}

type BrandsRepository struct {
	conn *sqlx.DB
}

func NewBrandsRepository(conn *sqlx.DB) store.BrandsRepository {
	return &BrandsRepository{conn: conn}
}

func (r BrandsRepository) Create(ctx context.Context, brand *models.Brand) error {
	_, err := r.conn.Exec("INSERT INTO brands(name) VALUES ($1)", brand.Name)
	if err != nil {
		return err
	}

	return nil
}

func (r BrandsRepository) All(ctx context.Context) ([]*models.Brand, error) {
	brands := make([]*models.Brand, 0)
	query := "SELECT * FROM brands"

	if err := r.conn.Select(&brands, query); err != nil {
		return nil, err
	}

	return brands, nil
}

func (r BrandsRepository) ByID(ctx context.Context, id uint) (*models.Brand, error) {
	brand := new(models.Brand)
	if err := r.conn.Get(brand, "SELECT * FROM brands WHERE id=$1", id); err != nil {
		return nil, err
	}

	return brand, nil
}

func (r BrandsRepository) Update(ctx context.Context, brand *models.Brand, id uint) error {
	_, err := r.conn.Exec("UPDATE brands SET name = $1 WHERE id = $2", brand.Name, id)
	if err != nil {
		return err
	}

	return nil
}

func (r BrandsRepository) Delete(ctx context.Context, id uint) error {
	_, err := r.conn.Exec("DELETE FROM brands WHERE id = $1", id)
	if err != nil {
		return err
	}

	return nil
}
