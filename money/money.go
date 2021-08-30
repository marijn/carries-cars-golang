package money

type CurrencyIsoCode string

type Money struct {
	amount          int
	currencyIsoCode CurrencyIsoCode
}

func EUR(amount int) Money {
	return Money{
		amount:          amount,
		currencyIsoCode: Euro,
	}
}

func USD(amount int) Money {
	return Money{
		amount:          amount,
		currencyIsoCode: UnitedStatesDollar,
	}
}

func (money Money) Equals(other Money) bool {
	return money.amount == other.amount && money.currencyIsoCode == other.currencyIsoCode
}

func (money Money) Amount() int {
	return money.amount
}

const Euro = "EUR"
const UnitedStatesDollar = "USD"
