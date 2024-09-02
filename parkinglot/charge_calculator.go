package main

import "time"

type ChargeCalculator interface {
	CalculateFees(entryTime time.Time, vehicleType VehicleType) int
}

type ChargeCalculatorImpl struct {
	PriceMap  map[VehicleType]Charge
	BaseHours float64
}

func (f ChargeCalculatorImpl) CalculateFees(entryTime time.Time, vehicleType VehicleType) int {
	currentTime := time.Now()
	timeDifference := currentTime.Sub(entryTime)
	charge := f.PriceMap[vehicleType]
	if timeDifference.Hours() < f.BaseHours {
		return charge.Base
	}
	minutesInterval := int(timeDifference.Minutes() / 30)
	if int(timeDifference.Minutes())%30 != 0 {
		minutesInterval += 1
	}
	return minutesInterval*charge.Extra + charge.Base
}
