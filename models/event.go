package models

import (
	"fmt"
)

type BaseEvent struct {
	Id             int64     `json:"id"`
	Title          string    `json:"Title"`
	Subtitle       string    `json:"Subtitle"`
	LocationName   string    `json:"LocationName"`
	LocationCoords []float64 `json:"LocationCoords"`
	Capacity       int64     `json:"Capacity"`
	Bookings       int64     `json:"Bookings"`
	Sponsored      bool      `json:"Sponsored"`
	Tags           string    `json:"Tags"`
}

type Event interface {
	GetId() int64
	GetTitle() string
	GetSubtitle() string
	GetLocationName() string
	GetLocationCoords() []float64
	GetCapacity() int64
	GetBookings() int64
	GetSponsored() bool
	GetPercentBooked() string
	GetTags() string
	ToString() string
}

func (event BaseEvent) GetId() int64 {
	return event.Id
}

func (event BaseEvent) GetTitle() string {
	return event.Title
}

func (event BaseEvent) GetSubtitle() string {
	return event.Subtitle
}

func (event BaseEvent) GetBookings() int64 {
	return event.Bookings
}

func (event BaseEvent) GetLocationName() string {
	return event.LocationName
}

func (event BaseEvent) GetLocationCoords() []float64 {
	return event.LocationCoords
}

func (event BaseEvent) GetCapacity() int64 {
	return event.Capacity
}

func (event BaseEvent) GetTags() string {
	return event.Tags
}

func (event BaseEvent) GetSponsored() bool {
	return event.Sponsored
}

func (event BaseEvent) GetPercentBooked() string {
	return fmt.Sprintf("%.2f%%", (float64(event.Bookings)/float64(event.Capacity))*100)
}

func (event BaseEvent) ToString() string {
	return fmt.Sprintf(
		"%s: %s \n====================== \nLocation: %s \nCoordinates: (%f, %f) \nCapacity: %d \nBookings: %d \nTags: %s\n\n",
		event.GetTitle(),
		event.GetSubtitle(),
		event.GetLocationName(),
		event.GetLocationCoords()[0],
		event.GetLocationCoords()[1],
		event.GetCapacity(),
		event.GetBookings(),
		event.GetTags(),
	)
}

func InitEvent(Id int64, Title string, Subtitle string, LocationName string, LocationCoords []float64, Capacity int64, Bookings int64, Sponsored bool, Tags string) BaseEvent {
	event := new(BaseEvent)
	event.Id = Id
	event.Title = Title
	event.Subtitle = Subtitle
	event.LocationName = LocationName
	event.LocationCoords = LocationCoords
	event.Capacity = Capacity
	event.Bookings = Bookings
	event.Sponsored = Sponsored
	event.Tags = Tags
	return *event
}
