package main

import "time"

type Genre int

const (
	Horror Genre = iota
	Romance
	Comedy
	Action
	SciFi
)

type Language int

const (
	Hindi Language = iota
	English
	French
)

type Movie struct {
	Id          string
	Name        string
	ReleaseData time.Time
	CastList    []Actor
	Language
	Genre
}

type Actor struct {
	Id   string
	Name string
}

type Cinema struct {
	Id     string
	Name   string
	CityId string
}

type Show struct {
	Id       string
	MovieId  string
	CinemaId string
}

type City struct {
	Id   string
	Name string
}

type Seat struct {
	Id       string
	ShowId   string
	IsBooked bool
}

type Booking struct {
	Id          string
	UserId      string
	SeatId      string
	BookingTime time.Time
	PaymentId   string
}

type SearchQuery struct {
	Title    *string
	Language *Language
	Genre    *Genre
	City     string
}
