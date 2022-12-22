package models

import "time"

type Trips struct {
	TripID          uint   `json:"id" gorm:"primaryKey"`
	DestinationName string `json:"destinationName" gorm:"type: varchar(255)"`
	Title           string
	CreatorID       int
	Accomodation    string
	Transportation  string
	Eatenary        string
	FromDate        time.Time
	ToDate          time.Time
	Description     string
	Transactions    []Transactions `gorm:"foreignKey:TripID"`
	ImageTrips      []ImageTrips   `gorm:"foreignKey:TripID"`
	Country         []Countries    `gorm:"foreignKey:IDCountries"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	Price           uint
	Quantity        uint
}
