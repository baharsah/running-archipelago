package repo

import (
	"baharsah/models"

	"gorm.io/gorm"
)

type UserRepo interface {
	GetUsers() ([]models.User, error)
	GetUser(models.User) (models.User, error)
	SetUser(models.User) (models.User, error)
}

type repo struct {
	db *gorm.DB
}

func RepoUser(db *gorm.DB) *repo {
	return &repo{db}
}

func (r *repo) GetUsers() ([]models.User, error) {
	var users []models.User
	err := r.db.Find(&users).Error

	return users, err
}

func (r *repo) SetUser(user models.User) (models.User, error) {

	err := r.db.Create(&user).Error

	return user, err
}

func (r *repo) GetUser(data models.User) (models.User, error) {
	var user models.User
	err := r.db.Where("email = ? ", data.Email).First(&user).Error
	return user, err
}
