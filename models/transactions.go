package models

type Transactions struct {
	ID            int `gorm:"primaryKey"`
	UserID        uint
	TripID        uint
	Trips         Trips `gorm:"foreignKey:TripID"`
	TransferProof string
	PaymentStatus int
	ApprovalID    int
	Quantity      uint
}

// disini dikirim keluar
