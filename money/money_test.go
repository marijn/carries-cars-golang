package money_test

import (
	money "carries-cars.com/money"
	"testing"
)

// This test is only here to verify that false test results get detected properly in GitHub Actions
func Test_Fails(t *testing.T) {
	t.Fatalf("Make it red on GitHub Actions (%s)", money.Euro)
}

// This test is only here to verify that false test results get detected properly in GitHub Actions
func Test_Money_Exposes_Currency_IsoCode_for_Euro(t *testing.T) {
	if money.Euro != "EUR" {
		t.Fatalf("Currency IsoCode of Euro not exposed, expected = 'EUR', actual = '%s'", money.Euro)
	}
}

func Test_Money_Equals_detects_equal_values(t *testing.T) {
	t.Skip("Todo")
}

func Test_Money_Equals_detects_currency_differences(t *testing.T) {
	t.Skip("Todo")
}

func Test_Money_Equals_detects_amount_differences(t *testing.T) {
	t.Skip("Todo")
}
