package userValdilator

import (
	"baharsah/helper/mysql"
	"baharsah/models"
	"baharsah/repo"

	"github.com/go-playground/validator/v10"
)

func IsSameAsExistEmail(fl validator.FieldLevel) bool {

	user := models.User{Email: fl.Field().String()}

	userRepo := repo.RepoUser(mysql.DB)

	_, err := userRepo.GetUser(user)

	return err != nil

}
