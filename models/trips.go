package models

import "time"

type Trips struct {
	ID              int
	DestinationName string `json:"destinationName" gorm:"type: varchar(255)"`
	Title           string
	Accomodation    string
	Transportation  string
	Eatenary        string
	FromDate        time.Time
	ToDate          time.Time
	Description     string
	ImageID         int
	ImageTrips      []ImageTrips `gorm:"many2many:trip_image;onDelete:CASCADE;onUpdate:CASCADE"`
	// referencing to CtryID
	//preloading the Countries struct
	Country   Countries `gorm:"foreignKey:CtryID;onDelete:CASCADE;onUpdate:CASCADE"`
	CtryID    uint
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Price     uint
	Quantity  uint
}
