package database

import (
	"github.com/wduartebr/goexpert/apis/internal/entity"
)

type UserInterface interface {
	CreateUser(user *entity.User) error
	FindByEmail(email string) (*entity.User, error)
}

type ProductInterface interface {
	CreateProduct(product *entity.Product) error
	FindById(id string) (*entity.Product, error)
	FindAll(page, limit int, sort string) ([]entity.Product, error)
	UpdateProduct(product *entity.Product) error
	DeleteProduct(id string) error
}
