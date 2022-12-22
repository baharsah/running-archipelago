package models

type Transactions struct {
	ID            uint `gorm:"primaryKey"`
	UserID        uint
	TripID        uint
	TransferProof string
	PaymentStatus int
	ApprovalID    int
	Quantity      uint
}
