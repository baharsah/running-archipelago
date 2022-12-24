package tripdito

type TripRequest struct {
	DestinationName string `form:"destination_name"`
	Title           string `form:"title"`
	CountryID       int    `form:"country_id"`
	Accomodation    string `form:"accomodation"`
	DateTrip        string `form:"date_trip"`
	Transportation  string `json:"transportation"`
	Eatenary        string `form:"eatenary"`
	// bagian ini akan dirubah menjadi sebuah tanggal yang di-"moment"-kan untuk menghitung jarak
	Day         int    `form:"day"`
	Night       int    `form:"night"`
	Quota       uint   `form:"quota"`
	Description string `form:"description"`
}
