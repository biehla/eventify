package database

import (
	"encoding/csv"
	"eventify/models"
	"fmt"
	"os"
	"strconv"
)

var events = make([]models.Event, 0, 300)

func setupEventDB(filename string) {
	type field int

	const (
		ID field = iota
		TITLE
		SUBTITLE
		LOCATION
		CAPACITY
		BOOKINGS
		SPONSORED
		TAGS
	)

	var fieldName = map[field]string{
		ID:        "ID",
		TITLE:     "Title",
		SUBTITLE:  "Subtitle",
		LOCATION:  "Location",
		CAPACITY:  "Capacity",
		BOOKINGS:  "Bookings",
		SPONSORED: "Sponsored",
		TAGS:      "Tags",
	}

	fileHandle, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file: ", err)
		return
	}
	defer fileHandle.Close()

	CSVReader := csv.NewReader(fileHandle)
	eventsArr, err := CSVReader.ReadAll()
	if err != nil {
		fmt.Println("Error reading csv: ", err)
		return
	}

	for eventID, event := range eventsArr {
		if eventID == 0 {
			continue
		}

		capacity, err := strconv.ParseInt(event[CAPACITY], 10, 64)
		if err != nil {
			fmt.Printf("Error parsing %s integer: %s\n", fieldName[CAPACITY], err)
		}

		bookings, err := strconv.ParseInt(event[BOOKINGS], 10, 64)
		if err != nil {
			fmt.Printf("Error parsing %s integer: %s\n", fieldName[BOOKINGS], err)
		}

		sponsored, err := strconv.ParseBool(event[SPONSORED])
		if err != nil {
			fmt.Printf("Error parsing %s integer: %s\n", fieldName[SPONSORED], err)
		}

		newEvent := new(models.Event)
		newEvent.Id = int64(eventID)
		newEvent.Title = event[TITLE]
		newEvent.Subtitle = event[SUBTITLE]
		newEvent.Capacity = capacity
		newEvent.Bookings = bookings
		newEvent.Sponsored = sponsored
		newEvent.Tags = event[TAGS]

		events = append(events, *newEvent)
	}
}
