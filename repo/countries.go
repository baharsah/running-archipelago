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

func (r *repo) GetCounties() ([]models.Countries, error) {

	var countries []models.Countries
	err := r.db.Debug().Find(&countries).Error
	return countries, err

}

func (r *repo) GetCountry(id int) (models.Countries, error) {
	var country models.Countries
	err := r.db.Debug().Find(&country, id).Error
	return country, err
}

func (r *repo) DeleteCountry(ctry models.Countries) (models.Countries, error) {
	err := r.db.Debug().Delete(&ctry).Error
	return ctry, err
}

func (r *repo) UpdateCountry(ctry models.Countries) (models.Countries, error) {
	err := r.db.Debug().Model(&ctry).Updates(&ctry).Error
	return ctry, err
}
