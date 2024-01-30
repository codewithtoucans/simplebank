package api

import "github.com/go-playground/validator/v10"

const (
	USD = "USD"
	EUR = "EUR"
	CAD = "CAD"
	RMB = "RMB"
)

var validCurrency validator.Func = func(fieldLevel validator.FieldLevel) bool {
	if currency, ok := fieldLevel.Field().Interface().(string); ok {
		return isSupportedCurrency(currency)
	}
	return false
}

func isSupportedCurrency(currency string) bool {
	switch currency {
	case USD, EUR, CAD, RMB:
		return true
	}
	return false
}
