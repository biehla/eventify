package models

type Event struct {
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
