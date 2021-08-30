package money_test

import (
	"carries-cars.com/money"
	"testing"
)

func Test_Money_Equals_detects_equal_values(t *testing.T) {
	actual := money.EUR(99).Equals(money.EUR(99))
	expected := true

	if actual != expected {
		t.Fatalf("EUR(99).Equals(EUR(99)) want = %t, have = %t", expected, actual)
	}
}

func Test_Money_Equals_detects_currency_differences(t *testing.T) {
	actual := money.EUR(10).Equals(money.USD(10))
	expected := false

	if actual != expected {
		t.Fatalf("EUR(10).Equals(USD(10)) want = %t, have = %t", expected, actual)
	}
}

func Test_Money_Equals_detects_amount_differences(t *testing.T) {
	actual := money.EUR(1).Equals(money.EUR(2))
	expected := false

	if actual != expected {
		t.Fatalf("EUR(1).Equals(EUR(2)) want = %t, have = %t", expected, actual)
	}
}

func Test_Money_Multiply_multiplies(t *testing.T) {
	actual := money.EUR(200).Multiply(2.00)
	expected := money.EUR(400)

	if actual != expected {
		t.Fatalf("EUR(200).Multiply(2.00) want = EUR(%v), have = EUR(%v)", expected.Amount(), actual.Amount())
	}
}

func Test_Money_Multiply_rounds_upward_correctly(t *testing.T) {
	actual := money.EUR(100).Multiply(1.999)
	expected := money.EUR(200)

	if actual != expected {
		t.Fatalf("EUR(100).Multiply(1.999) want = EUR(%v), have = EUR(%v)", expected.Amount(), actual.Amount())
	}
}

func Test_Money_Multiply_rounds_downward_correctly(t *testing.T) {
	actual := money.EUR(100).Multiply(1.994)
	expected := money.EUR(199)

	if actual != expected {
		t.Fatalf("EUR(100).Multiply(1.994) want = EUR(%v), have = EUR(%v)", expected.Amount(), actual.Amount())
	}
}

func Test_Money_Amount_exposes_value(t *testing.T) {
	t.Skip("Todo")
}

func Test_Money_CurrencyIsoCode_exposes_value(t *testing.T) {
	t.Skip("Todo")
}
