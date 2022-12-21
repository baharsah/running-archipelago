package repo

import (
	"baharsah/models"

	"gorm.io/gorm"
)

type UserRepo interface {
	GetUsers() ([]models.User, error)
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
