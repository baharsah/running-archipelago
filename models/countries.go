package models

type Countries struct {
	IDCountries uint `gorm:"primaryKey"`
	Country     string
}
