package database

import (
	"encoding/csv"
	"eventify/models"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type eventsArray []models.Event

var events = make(eventsArray, 0, 300)

type EventsDB interface {
	GetEvent(id int) models.Event
	SetEvent(id int, event models.Event) bool
}

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
	defer func(fileHandle *os.File) {
		err := fileHandle.Close()
		if err != nil {
			fmt.Println("Error closing file: ", err)
		}
	}(fileHandle)

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

		re, err := regexp.Compile(`(\w*)\s\((\d*\.\d*).*?,\s(\d*\.\d*)`)
		if err != nil {
			fmt.Println("Error compiling regex: ", err)
			panic(err)
		}

		matches := re.FindStringSubmatch(event[LOCATION])
		if len(matches) < 3 {
			fmt.Println("Error parsing location: ", err)
			continue
		}
		latitude, err := strconv.ParseFloat(matches[2], 64)
		if err != nil {
			fmt.Println("Error parsing latitude: ", err)
			panic(err)
		}

		longitude, err := strconv.ParseFloat(matches[3], 64)
		if err != nil {
			fmt.Println("Error parsing longitude: ", err)
			panic(err)
		}

		coordinates := []float64{latitude, longitude}

		newEvent := new(models.BaseEvent)
		newEvent.Id = int64(eventID)
		newEvent.Title = event[TITLE]
		newEvent.Subtitle = event[SUBTITLE]
		newEvent.LocationName = matches[1]
		newEvent.LocationCoords = coordinates
		newEvent.Capacity = capacity
		newEvent.Bookings = bookings
		newEvent.Sponsored = sponsored
		newEvent.Tags = event[TAGS]

		events = append(events, newEvent)
	}
}

func (events eventsArray) GetEvent(id int64) models.Event {
	return events[id]
}

func (events eventsArray) SetEvent(id int, newEvent models.Event) bool {
	events[id] = newEvent
	return true // TODO: at some point make this do some validation or write a validation function
}
