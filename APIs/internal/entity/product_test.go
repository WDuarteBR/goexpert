package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewProduct(t *testing.T) {
	p, err := NewProduct("prd one", 100.0)

	assert.NotNil(t, p)
	assert.Nil(t, err)
	assert.NotEmpty(t, p.ID)
	assert.Equal(t, "prd one", p.Name)
	assert.Equal(t, 100.0, p.Price)
}

func TestProductWhenNameIsRequired(t *testing.T) {
	p, err := NewProduct("", 100.0)

	assert.Nil(t, p)
	assert.Equal(t, ErrNameIsRequired, err)
}

func TestProductWhenPriceIsRequired(t *testing.T) {
	p, err := NewProduct("prd one", 0.0)

	assert.Nil(t, p)
	assert.Equal(t, ErrPriceIsReqired, err)
}

func TestProductWhenPirceInvalid(t *testing.T) {
	p, err := NewProduct("prd one", -10.0)

	assert.Nil(t, p)
	assert.Equal(t, ErrInvalidPrice, err)

}

func TestProductValidate(t *testing.T) {
	p, err := NewProduct("prd one", 1000.0)

	assert.Nil(t, err)
	assert.NotNil(t, p)
	assert.Nil(t, p.Validate())
}
