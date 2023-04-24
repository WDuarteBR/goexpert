//go:build wireinject
// +build wireinject

// go:build wireinject
package main

import (
	"database/sql"
	"github.com/google/wire"
	"github.com/wduartebr/goexpert/di/product"
)

var setRepositoryDependency = wire.NewSet(
	product.NewProductRepository,
	wire.Bind(new(product.ProductRepositoryInterface), new(*product.ProductRepository))
)
func NewUseCaseByWire(db *sql.DB) *product.ProductUseCase {
	wire.Build(
		setRepositoryDependency,
		product.NewProductUseCase,
	)
	return &product.ProductUseCase{}
}
