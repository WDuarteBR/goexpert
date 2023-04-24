package product

import "database/sql"

type ProductRepositoryInterface interface {
	GetProduct(id int) (*Product, error)
}

type ProductRepository struct {
	Db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{
		Db: db,
	}
}

func (p *ProductRepository) GetProduct(id int) (Product, error) {
	return Product{
		Id:   id,
		Name: "Product fake",
	}, nil
}
