package main

type BookingService struct {
	store *Store
}

func (b BookingService) GetBookings(userId string) []Booking {
	bookings := make([]Booking, 0)
	for _, booking := range b.store.GetBookings() {
		if booking.UserId == userId {
			bookings = append(bookings, booking)
		}
	}
	return bookings
}

func (b BookingService) AddBooking(booking Booking) {
	b.store.AddBooking(booking)
}
