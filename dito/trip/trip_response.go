package tripdito

type TripResponse struct {
	Title          string   `json:"title"`
	CountryID      string   `json:"country_id"`
	Accomodation   string   `json:"accomodation"`
	Transportation string   `json:"transportation"`
	Eatenary       string   `json:"eat"`
	Day            string   `json:"day"`
	Night          string   `json:"night"`
	DateTrip       string   `json:"datetrip"`
	Price          string   `json:"price"`
	Quota          int      `json:"quota"`
	Description    string   `json:"description"`
	Image          []string `json:"image"`
}
