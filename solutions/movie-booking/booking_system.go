package main

import (
	"fmt"
	"time"
)

type BookingSystem interface {
	GetMovieDetails(movieId string) Movie
	SearchShows(query SearchQuery) []Show
	GetSeatMap(showId string) []Seat
	BookSeat(userId string, seatId string) (Booking, error)
	GetBookings(userId string) []Booking
}

type BookingSystemImpl struct {
	searchService  SearchService
	seatService    SeatService
	paymentService PaymentService
	bookingService BookingService
}

func (b BookingSystemImpl) GetMovieDetails(movieId string) Movie {
	//TODO implement me
	panic("implement me")
}

func (b BookingSystemImpl) SearchShows(query SearchQuery) []Show {
	return b.searchService.Search(query)
}

func (b BookingSystemImpl) GetSeatMap(showId string) []Seat {
	return b.seatService.GetSeatMap(showId)
}

func (b BookingSystemImpl) BookSeat(userId string, seatId string) (Booking, error) {
	booking := Booking{}
	reserveStatus := b.seatService.ReserveSeat(userId, seatId)
	if !reserveStatus {
		return booking, fmt.Errorf("seat reservation failed")
	}
	paymentId, err := b.paymentService.ProcessPayment(100, userId)
	if err != nil {
		return booking, err
	}
	err = b.seatService.BookSeat(seatId)
	if err != nil {
		return booking, err
	}
	booking.SeatId = seatId
	booking.BookingTime = time.Now()
	booking.Id = fmt.Sprintf("%s-%s", userId, seatId)
	booking.PaymentId = paymentId
	b.bookingService.AddBooking(booking)
	return booking, err
}

func (b BookingSystemImpl) GetBookings(userId string) []Booking {
	return b.bookingService.GetBookings(userId)
}
