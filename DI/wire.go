//go:build wireinject
// +build wireinject

// go:build wireinject
package main

import (
	"database/sql"
	"github.com/google/wire"
	"github.com/wduartebr/goexpert/di/product"
)

func NewUseCaseByWire(db *sql.DB) *product.ProductUseCase {
	wire.Build(
		product.NewProductRepository,
		product.NewProductUseCase,
	)
	return &product.ProductUseCase{}
}
