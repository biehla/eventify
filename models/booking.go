package models

type bookingType int64

const (
	EVENT_BOOKING bookingType = iota
	BUNDLED_BOOKING
)

var BookingType = map[bookingType]string{
	EVENT_BOOKING:   "event",
	BUNDLED_BOOKING: "bundledEvent",
}

type baseBooking struct {
	Id              int64       `json:"Id"`
	UserIdPublicKey int64       `json:"UserIdPublicKey"`
	GroupSize       int64       `json:"GroupSize"`
	BookingType     bookingType `json:"BookingType"`
}

type EventBooking struct {
	*baseBooking
	EventId int64 `json:"EventId"`
}

type BundledBooking struct {
	*baseBooking
	EventIds []int64 `json:"EventIds"`
}

type Booking interface {
	GetId() int64
	GetEventIds() []int64
	GetGroupSize() int64
	GetUserID() int64
	GetBookingType() bookingType
}

func (booking baseBooking) GetId() int64 {
	return booking.Id
}

func (booking baseBooking) GetGroupSize() int64 {
	return booking.GroupSize
}

func (booking baseBooking) GetUserID() int64 {
	return booking.UserIdPublicKey
}

func (booking baseBooking) GetBookingType() bookingType {
	switch booking.BookingType {
	case EVENT_BOOKING:
		return EVENT_BOOKING
	case BUNDLED_BOOKING:
		return BUNDLED_BOOKING
	}
	return -1
}

func InitEventBooking(Id int64, UserIdPublicKey int64, GroupSize int64, EventId int64) EventBooking {
	booking := new(EventBooking)
	booking.Id = Id
	booking.UserIdPublicKey = UserIdPublicKey
	booking.GroupSize = GroupSize
	booking.BookingType = EVENT_BOOKING
	booking.EventId = EventId
	return *booking
}

func InitBundledBooking(Id int64, UserIdPublicKey int64, GroupSize int64, EventIds []int64) BundledBooking {
	booking := new(BundledBooking)
	booking.Id = 1 // Should be Id TODO: Switch back
	booking.UserIdPublicKey = UserIdPublicKey
	booking.GroupSize = GroupSize
	booking.BookingType = EVENT_BOOKING
	booking.EventIds = EventIds
	return *booking
}

func (booking EventBooking) GetEventIds() []int64 {
	var EventId = make([]int64, 0, 1)
	EventId = append(EventId, booking.EventId)
	return EventId
}

func (booking BundledBooking) GetEventIds() []int64 {
	return booking.EventIds
}
