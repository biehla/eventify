package models

import (
	"fmt"
	"strings"
)

type bookingType int64

const (
	EVENT_BOOKING bookingType = iota
	BUNDLED_BOOKING
)

var BookingType = map[bookingType]string{
	EVENT_BOOKING:   "event",
	BUNDLED_BOOKING: "bundledEvent",
}

type booking struct {
	Id              int64       `json:"Id"`
	UserIdPublicKey int64       `json:"UserIdPublicKey"`
	GroupSize       int64       `json:"GroupSize"`
	BookingType     bookingType `json:"BookingType"`
	EventIds        []int64     `json:"EventIds"`
}

type Booking interface {
	GetId() int64
	GetGroupSize() int64
	GetUserID() int64
	GetBookingType() string
	ToString() string
	GetEventIds() []int64
}

func (booking booking) GetId() int64 {
	return booking.Id
}

func (booking booking) GetGroupSize() int64 {
	return booking.GroupSize
}

func (booking booking) GetUserID() int64 {
	return booking.UserIdPublicKey
}

func (booking booking) GetBookingType() string {
	return BookingType[booking.BookingType]
}

func (booking booking) ToString() string {
	eventBuilder := new(strings.Builder)
	var eventsLabel string

	if booking.GetBookingType() == BookingType[BUNDLED_BOOKING] {
		eventBuilder.WriteString("[")
		for _, id := range booking.GetEventIds() {
			str := fmt.Sprintf("%d, ", id)
			eventBuilder.WriteString(str)
		}
		eventBuilder.WriteString("]")
		eventsLabel = "Events"
	} else {
		eventBuilder.WriteString(fmt.Sprintf("%d", booking.GetEventIds()))
		eventsLabel = "Event"
	}

	return fmt.Sprintf(
		"Booking %d \n============== \nUser ID PubKey: %d \nGroup Size: %d\nBooking Type: %s \n%s: %s",
		booking.GetId(),
		booking.GetUserID(),
		booking.GetGroupSize(),
		booking.GetBookingType(),
		eventsLabel,
		eventBuilder.String(),
	)
}

func InitBooking(Id int64, UserIdPublicKey int64, GroupSize int64, bookingType bookingType, EventIds []int64) booking {
	booking := new(booking)
	booking.Id = Id
	booking.UserIdPublicKey = UserIdPublicKey
	booking.GroupSize = GroupSize
	booking.BookingType = bookingType
	booking.EventIds = EventIds
	return *booking
}

func (booking booking) GetEventIds() []int64 {
	return booking.EventIds
}
