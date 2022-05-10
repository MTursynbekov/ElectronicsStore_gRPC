package models

type ProductImage struct {
	ID        uint   `db:"id"`
	Src       string `db:"src"`
	ProductId uint   `db:"product_id"`
}
