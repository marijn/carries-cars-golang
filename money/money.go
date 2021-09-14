package money

import "math"

type CurrencyIsoCode string

type PublicMoney interface {
	Amount() int
	CurrencyIsoCode() CurrencyIsoCode
	MultiplyAndRound(multiplier float64) PublicMoney
	Equals(other PublicMoney) bool
}

type verifiedMoney struct {
	// amount is denoted in the lowest denominator of the corresponding currency.
	// E.g. amount is in whole cents for the Euro or UnitedStatesDollar
	amount          int
	currencyIsoCode CurrencyIsoCode
}

// EUR acts as a named constructor function to create verifiedMoney for the Euro currency.
// Provide the amount in cents.
func EUR(amount int) PublicMoney {
	return verifiedMoney{
		amount:          amount,
		currencyIsoCode: Euro,
	}
}

// USD acts as a named constructor function to create verifiedMoney for the UnitedStatesDollar currency.
// Provide the amount in cents.
func USD(amount int) PublicMoney {
	return verifiedMoney{
		amount:          amount,
		currencyIsoCode: UnitedStatesDollar,
	}
}

func (money verifiedMoney) Equals(other PublicMoney) bool {
	return money.amount == other.Amount() && money.currencyIsoCode == other.CurrencyIsoCode()
}

func (money verifiedMoney) CurrencyIsoCode() CurrencyIsoCode {
	return money.currencyIsoCode
}

func (money verifiedMoney) Amount() int {
	return money.amount
}

func (money verifiedMoney) MultiplyAndRound(multiplier float64) PublicMoney {
	multipliedAmount := float64(money.amount) * multiplier
	multipliedAmountRounded := int(math.Round(multipliedAmount))

	return verifiedMoney{
		amount:          multipliedAmountRounded,
		currencyIsoCode: money.currencyIsoCode,
	}
}

const Euro = "EUR"
const UnitedStatesDollar = "USD"
