package database

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/wduartebr/goexpert/apis/internal/entity"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestCreateProduct(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&entity.Product{})
	product, err := entity.NewProduct("Prd one", 100.00)
	assert.NoError(t, err)
	productDB := NewProduct(db)
	err = productDB.CreateProduct(product)
	assert.Nil(t, err)
}

func TestFindAll(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&entity.Product{})
	for i := 1; i < 24; i++ {
		p, err := entity.NewProduct(fmt.Sprintf("Produto %d", i), rand.Float64()*100)
		assert.Nil(t, err)
		db.Create(p)
	}
	productDB := NewProduct(db)

	products, err := productDB.FindAll(1, 10, "asc")
	assert.NoError(t, err)
	assert.Len(t, products, 10)
	assert.Equal(t, "Produto 1", products[0].Name)
	assert.Equal(t, "Produto 10", products[9].Name)

	products, err = productDB.FindAll(2, 10, "asc")
	assert.NoError(t, err)
	assert.Len(t, products, 10)
	assert.Equal(t, "Produto 11", products[0].Name)
	assert.Equal(t, "Produto 20", products[9].Name)

	products, err = productDB.FindAll(3, 10, "asc")
	assert.NoError(t, err)
	assert.Len(t, products, 3)
	assert.Equal(t, "Produto 21", products[0].Name)
	assert.Equal(t, "Produto 23", products[2].Name)

}
