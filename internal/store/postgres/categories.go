package postgres

import (
	"context"

	"github.com/MTursynbekov/ElectronicsStore_gRPC/internal/models"
	"github.com/MTursynbekov/ElectronicsStore_gRPC/internal/store"
	"github.com/jmoiron/sqlx"
)

func (db *DB) Categories() store.CategoriesRepository {
	if db.categories == nil {
		db.categories = NewCategoriesRepository(db.conn)
	}

	return db.categories
}

type CategoriesRepository struct {
	conn *sqlx.DB
}

func NewCategoriesRepository(conn *sqlx.DB) store.CategoriesRepository {
	return &CategoriesRepository{conn: conn}
}

func (r CategoriesRepository) Create(ctx context.Context, category *models.Category) error {
	_, err := r.conn.Exec("INSERT INTO categories(name) VALUES ($1)", category.Name)
	if err != nil {
		return err
	}

	return nil
}

func (r CategoriesRepository) All(ctx context.Context) ([]*models.Category, error) {
	categories := make([]*models.Category, 0)
	query := "SELECT * FROM categories"

	if err := r.conn.Select(&categories, query); err != nil {
		return nil, err
	}

	return categories, nil
}

func (r CategoriesRepository) ByID(ctx context.Context, id uint) (*models.Category, error) {
	category := new(models.Category)
	if err := r.conn.Get(category, "SELECT * FROM categories WHERE id=$1", id); err != nil {
		return nil, err
	}

	return category, nil
}

func (r CategoriesRepository) Update(ctx context.Context, category *models.Category, id uint) error {
	_, err := r.conn.Exec("UPDATE categories SET name = $1 WHERE id = $2", category.Name, id)
	if err != nil {
		return err
	}

	return nil
}

func (r CategoriesRepository) Delete(ctx context.Context, id uint) error {
	_, err := r.conn.Exec("DELETE FROM categories WHERE id = $1", id)
	if err != nil {
		return err
	}

	return nil
}
