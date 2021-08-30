package pricingEngine

import (
	"carries-cars.com/money"
)

func CalculatePrice(pricePerMinute money.Money, duration Duration) money.Money {
	return
}

type Duration struct {
	durationInMinutes int
}

func DurationInMinutes(durationInMinutes int) Duration {
	// TODO: Guard against durationInMinutes of zero or less

	return
}
