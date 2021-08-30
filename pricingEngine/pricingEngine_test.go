package pricingEngine_test

import (
	"carries-cars.com/money"
	pricingEngine "carries-cars.com/pricingEngine"
	"testing"
)

func Test_CalculatePrice_charged_per_minute(t *testing.T) {
	pricePerMinute := money.EUR(30)
	duration, _ := pricingEngine.DurationInMinutes(1)
	expected := money.EUR(30)

	if !pricingEngine.CalculatePrice(pricePerMinute, duration).Equals(expected) {
		t.Fatalf("Price EUR(30) x 1min, want = EUR(30), have = EUR(%v)", expected.Amount())
	}
}

func Test_Duration_guards_against_zero_or_negative_duration(t *testing.T) {
	_, err := pricingEngine.DurationInMinutes(0)
	expected := "Duration should be a positive number in minutes"

	if nil == err {
		t.Fatalf("DurationInMinutes(0), want = error(%q), have = nil", expected)
	}

	actual := err.Error()

	if expected != actual {
		t.Fatalf("DurationInMinutes(0), want = error(%q), have = error(%q)", expected, actual)
	}
}
