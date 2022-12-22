package models

type ImageTrips struct {
	ImageID uint   `json:"ImageID" gorm:"primaryKey"`
	URL     string `json:"URL"`
	TripID  uint   `json:"TripID"`
}
