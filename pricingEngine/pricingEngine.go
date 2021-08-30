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

// UnverifiedDuration should be used when accepting input from untrusted sources (pretty much anywhere) in the model.
// This type models input that has not been verified and is therefore unsafe to use until it has been verified.
// Use Verify() to transform it to trusted input in the form of a Duration model.
type UnverifiedDuration struct {
	DurationInMinutes int
}

func (unsafe UnverifiedDuration) Verify() (Duration, error) {
	return DurationInMinutes(unsafe.DurationInMinutes)
}

func DurationInMinutes(durationInMinutes int) (Duration, error) {
	if durationInMinutes <= 0 {
		defaultDuration := Duration{durationInMinutes: 1}

		return defaultDuration, errors.New("duration should be a positive number in minutes")
	}

	return Duration{durationInMinutes: durationInMinutes}, nil
}
