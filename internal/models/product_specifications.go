package models

type ProductSpecification struct {
	ID        uint   `db:"id"`
	Key       string `db:"key"`
	Value     string `db:"value"`
	ProductId uint   `db:"product_id"`
}
