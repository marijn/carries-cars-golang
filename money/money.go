package money

type CurrencyIsoCode string

type Money struct {
	amount          int
	currencyIsoCode CurrencyIsoCode
}

func Eur(amount int) Money {
	return
}

func (money Money) Equals(other Money) bool {
	return
}

func (money Money) Amount() int {
	return
}

const Euro = "EUR"
