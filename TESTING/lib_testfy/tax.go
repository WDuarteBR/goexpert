package tax

import "errors"

type Repository interface {
	SaveTax(amout float64) error
}

func CalculateTaxSave(amout float64, repository Repository) error {
	tax := CalculateTaxFE(amout)
	return repository.SaveTax(tax)

}

func CalculateTax(amout float64) (float64, error) {
	if amout <= 0 {
		return 0, errors.New("amout deve ser maior que zero")
	}
	if amout >= 1000 && amout < 20000 {
		return 10, nil
	}
	if amout >= 20000 {
		return 20.0, nil
	}

	return 5.0, nil
}

func CalculateTaxFE(amout float64) float64 {
	if amout <= 0 {
		return 0
	}
	if amout >= 1000 && amout < 20000 {
		return 10
	}
	if amout >= 20000 {
		return 20.0
	}

	return 5.0
}
