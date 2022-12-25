package transactionsdito

type TransactionsRequest struct {
	Qty        int `json:"counterQty"`
	Status     int `json:"status"`
	Attachment string
	TripID     uint `json:"tripId"`
	UserID     uint `json:"userId"`
}
