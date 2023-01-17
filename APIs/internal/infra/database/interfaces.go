package database

import (
	"github.com/wduartebr/goexpert/apis/internal/entity"
)

type UserInterface interface {
	CreateUser(user *entity.User) error
	FindByEmail(email string) (*entity.User, error)
}
