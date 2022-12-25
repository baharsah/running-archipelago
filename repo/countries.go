package repo

import (
	"baharsah/models"

	"gorm.io/gorm"
)

type CountryRepo interface {
	SetCountry(models.Countries) (models.Countries, error)
	GetCountry(id int) (models.Countries, error)
	GetCounties() ([]models.Countries, error)
	UpdateCountry(models.Countries) (models.Countries, error)
	DeleteCountry(models.Countries) (models.Countries, error)
}

func RepoCountry(db *gorm.DB) *repo {
	return &repo{db}
}

func (r *repo) SetCountry(ctry models.Countries) (models.Countries, error) {
	err := r.db.Debug().Create(&ctry).Error
	return ctry, err

}
