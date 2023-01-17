package entity

import (
	"errors"
	"time"

	"github.com/wduartebr/goexpert/apis/pkg/entity"
)

var (
	ErrIdRequired     = errors.New("id is required")
	ErrInvalidId      = errors.New("invalid id")
	ErrNameIsRequired = errors.New("name is required")
	ErrPriceIsReqired = errors.New("price is required")
	ErrInvalidPrice   = errors.New("invalid price")
)

type Product struct {
	ID        entity.ID `json:"id"`
	Name      string    `json:"name"`
	Price     int       `json:"price"`
	CreatedAt time.Time `json:"created_at"`
}

func NewProduct(name string, prince int) (*Product, error) {
	product := &Product{
		ID:        entity.NewId(),
		Name:      name,
		Price:     prince,
		CreatedAt: time.Now(),
	}

	err := product.Validate()
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (p *Product) Validate() error {
	if p.ID.String() == "" {
		return ErrIdRequired
	}

	if _, err := entity.ParseID(p.ID.String()); err != nil {
		return ErrInvalidId
	}

	if p.Name == "" {
		return ErrNameIsRequired
	}

	if p.Price == 0 {
		return ErrPriceIsReqired
	}

	if p.Price < 0 {
		return ErrInvalidPrice
	}

	return nil
}
