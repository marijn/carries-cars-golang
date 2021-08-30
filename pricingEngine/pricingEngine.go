package pricingEngine

import (
	"carries-cars.com/money"
)

func CalculatePrice(pricePerMinute money.Money, duration Duration) money.Money {
	return pricePerMinute.Multiply(float64(duration.durationInMinutes))
}

type Duration struct {
	durationInMinutes int
}

func DurationInMinutes(durationInMinutes int) (Duration, error) {
	return Duration{durationInMinutes: durationInMinutes}, nil
}
