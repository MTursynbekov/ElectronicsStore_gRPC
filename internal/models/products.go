package models

type Product struct {
	ID          uint    `db:"id"`
	Name        string  `db:"name"`
	Price       float32 `db:"price"`
	CategoryId  uint    `db:"category_id"`
	BrandId     uint    `db:"brand_id"`
	Description string  `db:"description"`
}
