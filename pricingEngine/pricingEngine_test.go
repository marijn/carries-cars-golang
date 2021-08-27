package pricingEngine_test

import (
	pricingEngine "carries-cars.com/pricingEngine"
	"testing"
)

// This test is only here to verify that false test results get detected properly in GitHub Actions
func Test_Fails(t *testing.T) {
	if pricingEngine.CalculatePrice() {
		t.Fatalf("Make it red on GitHub Actions")
	}
}
