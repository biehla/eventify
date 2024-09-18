package database

import (
	"encoding/csv"
	"eventify/models"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type bookingArray []models.Booking

var bookings = make(bookingArray, 0, 300)

type BookingDB interface {
	GetBooking(id int) models.Booking
	SetBooking(id int, booking models.Booking) bool
}

func setupBookingDB(filename string) {
	type field int

	const (
		ID field = iota
		USER_ID_PUBLIC_KEY
		GROUP_SIZE
		BOOKING_TYPE
		EVENT_ID
		EVENT_IDS
	)

	var fieldName = map[field]string{
		ID:                 "ID",
		USER_ID_PUBLIC_KEY: "User ID Public Key",
		GROUP_SIZE:         "Group Size",
		BOOKING_TYPE:       "Booking Type",
		EVENT_ID:           "Event ID",
		EVENT_IDS:          "Event IDs",
	}

	fileHandle, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file: ", err)
		return
	}
	defer func(fileHandle *os.File) {
		err := fileHandle.Close()
		if err != nil {
			fmt.Println("Error closing file: ", err)
		}
	}(fileHandle)

	CSVReader := csv.NewReader(fileHandle)
	bookingsArr, err := CSVReader.ReadAll()
	if err != nil {
		fmt.Println("Error reading csv: ", err)
		return
	}

	for bookingID, booking := range bookingsArr {
		var (
			eventIDsArr []int64
			newBooking  models.Booking
		)

		if bookingID == 0 {
			// fmt.Println("Continued!")
			continue
		}

		userIdPublicKey, err := strconv.ParseInt(booking[USER_ID_PUBLIC_KEY], 10, 64)
		if err != nil {
			fmt.Printf("Error parsing %s integer: %s\n", fieldName[USER_ID_PUBLIC_KEY], err)
		}

		groupSize, err := strconv.ParseInt(booking[GROUP_SIZE], 10, 64)
		if err != nil {
			fmt.Printf("Error parsing %s integer: %s\n", fieldName[GROUP_SIZE], err)
		}

		bookingType, err := strconv.ParseBool(booking[BOOKING_TYPE])
		if err != nil {
			fmt.Printf("Error parsing %s integer: %s\n", fieldName[BOOKING_TYPE], err)
		}

		if bookingType {
			eventIDsStringArr := strings.Split(booking[EVENT_IDS], ";")
			eventIDsArr = make([]int64, 0, len(eventIDsStringArr))
			for _, eventID := range eventIDsStringArr {
				eventID, err := strconv.ParseInt(eventID, 10, 64)
				if err != nil {
					fmt.Printf("Error parsing %s integer: %s\n", fieldName[EVENT_IDS], err)
				}
				eventIDsArr = append(eventIDsArr, int64(eventID))
			}
			newBooking = models.InitBooking(int64(bookingID), userIdPublicKey, groupSize, models.BUNDLED_BOOKING, eventIDsArr)
		} else {
			eventID, err := strconv.ParseInt(booking[EVENT_ID], 10, 64)
			if err != nil {
				fmt.Printf("Error parsing %s integer: %s\n", fieldName[EVENT_ID], err)
			}
			newBooking = models.InitBooking(int64(bookingID), userIdPublicKey, groupSize, models.EVENT_BOOKING, []int64{eventID})
		}

		bookings = append(bookings, newBooking)
	}
}

func (bookings bookingArray) GetBooking(id int64) models.Booking {
	return bookings[id]
}

func (bookings bookingArray) SetBooking(id int, newBooking models.Booking) bool {
	bookings[id] = newBooking
	return true // TODO: at some point make this do some validation or write a validation function
}
