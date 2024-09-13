package main

type ParkingStrategy interface {
	FindSlot(vehicleType VehicleType, floors []Floor) *Slot
}

type NearestAvailableStrategy struct {
}

func (n NearestAvailableStrategy) FindSlot(vehicleType VehicleType, floors []Floor) *Slot {
	for _, floor := range floors {
		slot := floor.GetSlot(vehicleType)
		if slot != nil {
			return slot
		}
	}
	return nil
}

type LeastOccupiedFloorStrategy struct {
}

func (l LeastOccupiedFloorStrategy) FindSlot(vehicleType VehicleType, floors []Floor) *Slot {
	var slot *Slot
	freeSlotCount := 0
	for _, floor := range floors {
		probableSpot := floor.GetSlot(vehicleType)
		if probableSpot != nil {
			if floor.FreeSlots > freeSlotCount {
				freeSlotCount = floor.FreeSlots
				slot = probableSpot
			}
		}
	}
	return slot
}
