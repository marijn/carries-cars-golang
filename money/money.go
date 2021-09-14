package money

import "math"

type CurrencyIsoCode string

type PublicMoney interface {
	Amount() int
	CurrencyIsoCode() CurrencyIsoCode
	MultiplyAndRound(multiplier float64) PublicMoney
	Equals(other PublicMoney) bool
}

type trustedMoney struct {
	// amount is denoted in the lowest denominator of the corresponding currency.
	// E.g. amount is in whole cents for the Euro or UnitedStatesDollar
	amount          int
	currencyIsoCode CurrencyIsoCode
}

// EUR acts as a named constructor function to create trustedMoney for the Euro currency.
// Provide the amount in cents.
func EUR(amount int) PublicMoney {
	return trustedMoney{
		amount:          amount,
		currencyIsoCode: Euro,
	}
}

// USD acts as a named constructor function to create trustedMoney for the UnitedStatesDollar currency.
// Provide the amount in cents.
func USD(amount int) PublicMoney {
	return trustedMoney{
		amount:          amount,
		currencyIsoCode: UnitedStatesDollar,
	}
}

func (money trustedMoney) Equals(other PublicMoney) bool {
	return money.amount == other.Amount() && money.currencyIsoCode == other.CurrencyIsoCode()
}

func (money trustedMoney) CurrencyIsoCode() CurrencyIsoCode {
	return money.currencyIsoCode
}

func (money trustedMoney) Amount() int {
	return money.amount
}

func (money trustedMoney) MultiplyAndRound(multiplier float64) PublicMoney {
	multipliedAmount := float64(money.amount) * multiplier
	multipliedAmountRounded := int(math.Round(multipliedAmount))

	return trustedMoney{
		amount:          multipliedAmountRounded,
		currencyIsoCode: money.currencyIsoCode,
	}
}

const Euro = "EUR"
const UnitedStatesDollar = "USD"
