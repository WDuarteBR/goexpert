package tax

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalculateTax(t *testing.T) {
	tax, err := CalculateTax(1000.0)
	assert.Nil(t, err)
	assert.Equal(t, 10.0, tax)

	tax, err = CalculateTax(0)
	assert.Error(t, err, "amout deve ser maior que zero")
	assert.Equal(t, 0.0, tax)
	assert.Contains(t, err.Error(), "ser maior")

}

func TestCalculateTaxSave(t *testing.T) {
	repository := &TaxRepositoryMock{}
	repository.On("SaveTax", 10.0).Return(nil)
	repository.On("SaveTax", 0.0).Return(errors.New("error saving tax"))

	err := CalculateTaxSave(1000.0, repository)
	assert.Nil(t, err)

	err = CalculateTaxSave(0.0, repository)
	assert.Error(t, err, "error saving tax")

	repository.AssertExpectations(t)

}
