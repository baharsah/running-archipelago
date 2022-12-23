package tripdito

import "time"

type TripRequest struct {
	Title          string
	CountryName    string
	Accomodation   string
	Transportation string
	Eatenary       string
	Day            int
	Night          int
	DateTrip       time.Time
	Price          uint
	Quota          uint
	Description    string
	// Image          string
	ImageURL []string
}
