package pricingEngine

import (
	"carries-cars.com/money"
	"errors"
)

func CalculatePrice(pricePerMinute money.Money, duration Duration) money.Money {
	return pricePerMinute.Multiply(float64(duration.durationInMinutes))
}

type Duration struct {
	durationInMinutes int
}

func DurationInMinutes(durationInMinutes int) (Duration, error) {
	if durationInMinutes <= 0 {
		return Duration{durationInMinutes: 1}, errors.New("Duration should be a positive number in minutes")
	}

	return Duration{durationInMinutes: durationInMinutes}, nil
}
