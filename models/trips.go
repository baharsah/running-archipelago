package models

type Trips struct {
	ID               int    `json:"id"`
	DestionationName string `json:"destinationName" gorm:"type: varchar(255)"`
}
