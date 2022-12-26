package transactionsdito

type TransactionsResponse struct {
	ID         int                  `json:"id"`
	CounterQty uint                 `json:"counterQty"`
	Status     int                  `json:"status"`
	Attachment string               `json:"attachment"`
	Trip       TripRelationResponse `json:"trip"`
}

type TripRelationResponse struct {
	ID           int    `json:"id"`
	Title        string `json:"title"`
	CountryID    int    `json:"country_id"`
	Accomodation string `json:"accomodation"`
	Eatenary     string `json:"eat"`
	Day          string `json:"day"`
	Night        string `json:"night"`
	DateTrip     string `json:"dateTrip"`
	Price        int    `json:"price"`
	Quota        int    `json:"quota"`
	Description  string `json:"description"`
}
