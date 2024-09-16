package database

import (
	"eventify/models"
)

const EVENT_FILE = "C:/Users/Anand/Documents/eventify/events.csv"
const BOOKING_FILE = "C:/Users/Anand/Documents/eventify/bookings.csv"

func GetEvent(id int) models.Event {
	return events[id]
}

func SetupDB() {
	setupEventDB(EVENT_FILE)
	setupBookingDB(BOOKING_FILE)
}
