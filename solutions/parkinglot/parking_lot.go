package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type ParkingLot struct {
	floors           []Floor
	parkingStrategy  ParkingStrategy
	tickets          map[string]Ticket
	chargeCalculator ChargeCalculator
}

func (p *ParkingLot) Park(vehicle Vehicle, entryTime time.Time) (Ticket, error) {
	slot := p.parkingStrategy.FindSlot(vehicle.Kind, p.floors)
	if slot == nil {
		return Ticket{}, fmt.Errorf("no slots available")
	}
	p.floors[slot.FloorId].UpdateSlotStatus(vehicle.Kind, slot.Id, true)
	slotId := fmt.Sprintf("%d-%d", slot.FloorId, slot.Id)
	ticket := Ticket{
		Id:        fmt.Sprintf("%s-%v", vehicle.Id, time.Now().Format("2006-01-02 15:04:05")),
		Vehicle:   vehicle,
		EntryTime: entryTime,
		SlotId:    slotId,
	}
	p.tickets[ticket.Id] = ticket
	return ticket, nil
}

func (p *ParkingLot) Remove(ticketId string) error {
	ticket, ok := p.tickets[ticketId]
	if !ok {
		return fmt.Errorf("ticket id does not exist")
	}
	tokens := strings.Split(ticket.SlotId, "-")
	floor, _ := strconv.Atoi(tokens[0])
	index, _ := strconv.Atoi(tokens[1])
	p.floors[floor].UpdateSlotStatus(ticket.Vehicle.Kind, index, false)
	return nil
}

func (p *ParkingLot) CalculateCharge(ticketId string) (int, error) {
	ticket, ok := p.tickets[ticketId]
	if !ok {
		return 0, fmt.Errorf("ticket id does not exist")
	}
	return p.chargeCalculator.CalculateFees(ticket.EntryTime, ticket.Vehicle.Kind), nil
}
