package main

import "sync"

type Store struct {
	movies       []Movie
	shows        []Show
	cinemas      []Cinema
	seats        []Seat
	bookings     []Booking
	seatMutex    sync.Mutex
	bookingMutex sync.Mutex
}

func (s *Store) GetMovieById(id string) *Movie {
	for _, movie := range s.movies {
		if movie.Id == id {
			return &movie
		}
	}
	return nil
}

func (s *Store) GetShows() []Show {
	return s.shows
}

func (s *Store) GetBookings() []Booking {
	return s.bookings
}

func (s *Store) GetCinemaById(id string) *Cinema {
	for _, cinema := range s.cinemas {
		if cinema.CityId == id {
			return &cinema
		}
	}
	return nil
}

func (s *Store) GetSeats() []Seat {
	return s.seats
}

func (s *Store) BookSeat(seatId string) bool {
	s.seatMutex.Lock()
	defer s.seatMutex.Unlock()
	for i := 0; i < len(s.seats); i += 1 {
		if s.seats[i].Id == seatId && !s.seats[i].IsBooked {
			s.seats[i].IsBooked = true
			return true
		}
	}
	return false
}

func (s *Store) AddBooking(booking Booking) {
	s.bookingMutex.Lock()
	defer s.bookingMutex.Unlock()
	s.bookings = append(s.bookings, booking)
}
