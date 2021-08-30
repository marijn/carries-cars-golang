package money

type CurrencyIsoCode string

type Money struct {
	amount          int
	currencyIsoCode CurrencyIsoCode
}

func EUR(amount int) Money {
	return
}

func USD(amount int) Money {
	return
}

func (money Money) Equals(other Money) bool {
	return
}

func (money Money) Amount() int {
	return
}

const Euro = "EUR"
const UnitedStatesDollar = "USD"
