package repo

import "baharsah/models"

type TripRepo interface {
	GetTrip() ([]models.Trips, error)
}

func (r *repo) GetTrip() ([]models.Trips, error) {

	var trips []models.Trips
	err := r.db.Find(&trips).Error
	return trips, err

}
