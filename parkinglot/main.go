package main

import (
	"fmt"
	"time"
)

func main() {
	chargeCalculator := ChargeCalculatorImpl{
		PriceMap: map[VehicleType]Charge{
			twoWheeler: {
				Base:  10,
				Extra: 2,
			},
			fourWheeler: {
				Base:  20,
				Extra: 4,
			},
		},
		BaseHours: 0,
	}
	parkingStrategy := NearestAvailableStrategy{}
	floorMap1 := map[VehicleType][]Slot{
		twoWheeler: {
			Slot{Id: 5, SupportedType: twoWheeler, FloorId: 0}},
		fourWheeler: {Slot{Id: 2, SupportedType: fourWheeler, FloorId: 0},
			Slot{Id: 4, SupportedType: fourWheeler, FloorId: 0},
			Slot{Id: 6, SupportedType: fourWheeler, FloorId: 0}},
	}
	floorMap2 := map[VehicleType][]Slot{
		fourWheeler: {Slot{Id: 2, SupportedType: fourWheeler, FloorId: 1},
			Slot{Id: 4, SupportedType: fourWheeler, FloorId: 1},
			Slot{Id: 6, SupportedType: fourWheeler, FloorId: 1}},
	}
	parkingLot := ParkingLot{
		floors: []Floor{{
			slotMap:   floorMap1,
			FreeSlots: 6,
		}, {
			slotMap:   floorMap2,
			FreeSlots: 6,
		}},
		parkingStrategy:  parkingStrategy,
		tickets:          make(map[string]Ticket),
		chargeCalculator: chargeCalculator,
	}

	vehicle1 := Vehicle{
		Id:   "id-1",
		Kind: twoWheeler,
	}
	vehicle2 := Vehicle{
		Id:   "id-2",
		Kind: twoWheeler,
	}
	ticket1, err := parkingLot.Park(vehicle1, time.Now().Add(-200*time.Minute))
	if err == nil {
		fmt.Println(ticket1)
		fee, _ := parkingLot.CalculateCharge(ticket1.Id)
		fmt.Println(fee)
	}
	ticket2, err := parkingLot.Park(vehicle2, time.Now().Add(-70*time.Minute))

	if err == nil {
		fmt.Println(ticket2)
		fee, _ := parkingLot.CalculateCharge(ticket2.Id)
		fmt.Println(fee)
	} else {
		fmt.Println(err)
	}
}
