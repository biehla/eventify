package database

import (
	"eventify/models"
)

const EVENT_FILE = "C:/Users/Anand/Documents/eventify/events.csv"
const BOOKING_FILE = "C:/Users/Anand/Documents/eventify/bookings.csv"

func GetEvent(id int64) models.Event {
	return events.GetEvent(id)
}

func SetEvent(event models.BaseEvent) bool {
	return events.SetEvent(int(event.Id), event)
}

func GetBooking(id int64) models.Booking {
	return bookings.GetBooking(id)
}

func SetBooking(booking models.Booking) bool {
	return bookings.SetBooking(int(booking.GetId()), booking)
}

func SetupDB() {
	setupEventDB(EVENT_FILE)
	setupBookingDB(BOOKING_FILE)
}
