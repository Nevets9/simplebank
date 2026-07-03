package util

// Constans for all supported currencies
const (
	USD = "USD"
	EUR = "EUR"
	IDR = "IDR"
)

// IsSupportedCurrency return true if the currency is supported
func IsSupportedCurrency (currency string) bool {
	switch currency {
	case USD, EUR, IDR :
		return true
	}
	return false
}
