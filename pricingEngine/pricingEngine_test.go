package pricingEngine_test

import (
	"carries-cars.com/money"
	pricingEngine "carries-cars.com/pricingEngine"
	"testing"
)

func Test_CalculatePrice_charged_per_minute(t *testing.T) {
	pricePerMinute := money.Eur(30)
	duration := pricingEngine.DurationInMinutes(1)
	expected := money.Eur(30)

	if !pricingEngine.CalculatePrice(pricePerMinute, duration).Equals(expected) {
		t.Fatalf("Price Eur(30) x 1min, want = Eur(30), have = Eur(%v)", expected.Amount())
	}
}
