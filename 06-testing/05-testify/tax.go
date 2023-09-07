package tax

import "errors"

type Repository interface {
	SaveTax(amount float64) error
}

func CalculateTax(value float64) (float64, error) {
	if value <= 0 {
		return 0.0, errors.New("amount must be greater than 0")
	} else if value >= 1000 && value < 20000 {
		return 10.0, nil
	} else if value >= 20000 {
		return 20.0, nil
	}
	return 5.0, nil
}

func CalculateTax2(value float64) float64 {
	if value <= 0 {
		return 0
	} else if value >= 1000 && value < 20000 {
		return 10.0
	} else if value >= 20000 {
		return 20.0
	}
	return 5.0
}

func CalculateTaxAndSave(amount float64, repository Repository) error {
	tax := CalculateTax2(amount)
	return repository.SaveTax(tax)
}