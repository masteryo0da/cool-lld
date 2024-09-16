package main

import (
	"fmt"
	"sync"
)

type SeatService struct {
	store      Store
	reserveMap map[string]string
	mutex      sync.Mutex
}

func (s *SeatService) GetSeatMap(showId string) []Seat {
	seats := make([]Seat, 0)
	for _, seat := range s.store.GetSeats() {
		if seat.ShowId == showId {
			seats = append(seats, seat)
		}
	}
	return seats
}

func (s *SeatService) ReserveSeat(userId string, seatId string) bool {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if _, ok := s.reserveMap[seatId]; ok {
		return false
	}
	s.reserveMap[seatId] = userId
	return true
}

func (s *SeatService) BookSeat(seatId string) error {
	status := s.store.BookSeat(seatId)
	if !status {
		return fmt.Errorf("seat not present or seat already booked")
	}
	return nil
}
