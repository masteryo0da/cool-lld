package main

type Floor struct {
	slotMap   map[VehicleType][]Slot
	FreeSlots int
}

func (f *Floor) UpdateSlotStatus(vehicleType VehicleType, id int, occupiedStatus bool) {
	slots, ok := f.slotMap[vehicleType]
	if ok {
		for index, slot := range slots {
			if slot.Id == id {
				f.slotMap[vehicleType][index].IsOccupied = occupiedStatus
				if occupiedStatus {
					f.FreeSlots += 1
				} else {
					f.FreeSlots -= 1
				}
			}
		}
	}
}

func (f *Floor) GetSlot(vehicleType VehicleType) *Slot {
	for _, value := range f.slotMap {
		for _, slot := range value {
			if !slot.IsOccupied && slot.SupportedType == vehicleType {
				return &slot
			}
		}
	}
	return nil
}
