package main

import "time"

type VehicleType int

const (
	twoWheeler VehicleType = iota
	fourWheeler
)

type Slot struct {
	Id            int
	SupportedType VehicleType
	IsOccupied    bool
	FloorId       int
}

type Vehicle struct {
	Id   string
	Kind VehicleType
}

type Ticket struct {
	Id string
	Vehicle
	EntryTime time.Time
	SlotId    string
}

type Charge struct {
	Base  int
	Extra int
}
