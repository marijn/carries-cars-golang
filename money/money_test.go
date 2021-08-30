package money_test

import (
	"carries-cars.com/money"
	"testing"
)

func Test_Money_Equals_detects_equal_values(t *testing.T) {
	actual := money.EUR(99).Equals(money.EUR(99))
	expected := true

	if actual != expected {
		t.Fatalf("EUR(99).Equals(EUR(99)) want = %b, have = %b", expected, actual)
	}
}

func Test_Money_Equals_detects_currency_differences(t *testing.T) {
	actual := money.EUR(10).Equals(money.USD(10))
	expected := false

	if actual != expected {
		t.Fatalf("EUR(10).Equals(USD(10)) want = %b, have = %b", expected, actual)
	}
}

func Test_Money_Equals_detects_amount_differences(t *testing.T) {
	actual := money.EUR(1).Equals(money.EUR(2))
	expected := false

	if actual != expected {
		t.Fatalf("EUR(1).Equals(EUR(2)) want = %b, have = %b", expected, actual)
	}
}

func Test_Money_Amount_exposes_value(t *testing.T) {
	t.Skip("Todo")
}

func Test_Money_CurrencyIsoCode_exposes_value(t *testing.T) {
	t.Skip("Todo")
}
