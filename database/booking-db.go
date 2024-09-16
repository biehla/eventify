package database

import (
	"encoding/csv"
	"eventify/models"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var bookings = make([]models.Booking, 0, 300)

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
	defer fileHandle.Close()

	CSVReader := csv.NewReader(fileHandle)
	bookingsArr, err := CSVReader.ReadAll()
	if err != nil {
		fmt.Println("Error reading csv: ", err)
		return
	}

	for bookingID, booking := range bookingsArr {
		fmt.Println("Booking: ", booking)
		var (
			eventIDsArr []int64
			newBooking  models.Booking
		)

		if bookingID == 0 {
			fmt.Println("Continued!")
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
			eventIDsArr = make([]int64, 0, len((eventIDsStringArr)))
			for _, eventID := range eventIDsStringArr {
				eventID, err := strconv.ParseInt(eventID, 10, 64)
				if err != nil {
					fmt.Printf("Error parsing %s integer: %s\n", fieldName[EVENT_IDS], err)
				}
				eventIDsArr = append(eventIDsArr, int64(eventID))
			}
			newBooking = models.InitBundledBooking(int64(bookingID), userIdPublicKey, groupSize, eventIDsArr)
		} else {
			eventID, err := strconv.ParseInt(booking[EVENT_ID], 10, 64)
			newBooking = new(models.EventBooking)
			if err != nil {
				fmt.Printf("Error parsing %s integer: %s\n", fieldName[EVENT_ID], err)
			}
			newBooking = models.InitEventBooking(int64(bookingID), userIdPublicKey, groupSize, eventID)
		}

		bookings = append(bookings, newBooking)
	}
}
