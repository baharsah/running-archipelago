package repo

import (
	"baharsah/models"

	"gorm.io/gorm"
)

type TripRepo interface {
	GetTrips() ([]models.Trips, error)
	CreateTrip(models.Trips) (models.Trips, error)
	DeleteTrip(trip models.Trips) (models.Trips, error)
	GetTrip(id int) (models.Trips, error)
}

func RepoTrip(db *gorm.DB) *repo {
	return &repo{db}
}
func (r *repo) GetTrips() ([]models.Trips, error) {

	var trips []models.Trips
	err := r.db.Debug().Preload("ImageTrips").Preload("Country").Find(&trips).Error
	return trips, err

}

func (r *repo) CreateTrip(trips models.Trips) (models.Trips, error) {

	err := r.db.Debug().Create(&trips).Error

	return trips, err

}

func (repo *repo) DeleteTrip(trip models.Trips) (models.Trips, error) {

	err := repo.db.Debug().Delete(&trip).Error
	return trip, err
}

func (r *repo) GetTrip(id int) (models.Trips, error) {

	var trips models.Trips
	err := r.db.Debug().Preload("ImageTrips").Preload("Country").First(&trips).Error
	return trips, err

}
