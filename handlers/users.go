package handlers

import (
	resultDito "baharsah/dito/result"
	usersdito "baharsah/dito/users"
	"baharsah/helper/bcrypt"
	"baharsah/models"
	"baharsah/repo"
	userValdilator "baharsah/valdilator/uservaldilator"
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type handler struct {
	UserRepo repo.UserRepo
}

func HandlerUser(UserRepo repo.UserRepo) *handler {
	return &handler{UserRepo}

}

func (h *handler) GetUsers(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	users, err := h.UserRepo.GetUsers()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())

	}

	w.WriteHeader(http.StatusOK)
	// response := ResultDito.SuccessResult{Code: http.StatusOK, Data: users}
	response := resultDito.SuccessResult{Code: http.StatusOK, Data: users}
	json.NewEncoder(w).Encode(response)
}

func (h *handler) CreateUser(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	request := new(usersdito.RegisterRequest)

	if err := json.NewDecoder(req.Body).Decode(&request); err != nil {
		res.WriteHeader(http.StatusBadRequest)
		response := resultDito.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(res).Encode(response)
		return
	}

	validation := validator.New()
	validation.RegisterValidation("email_exist", userValdilator.IsSameAsExistEmail)
	err := validation.Struct(request)

	if err != nil {

		res.WriteHeader(http.StatusBadRequest)
		response := resultDito.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}

		json.NewEncoder(res).Encode(response)
		return
	}

	password, err := bcrypt.HashingPassword(request.Password)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		response := resultDito.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(res).Encode(response)
	}

	user := models.User{
		Name:     request.FullName,
		Email:    request.Email,
		Password: password,
	}
	data, err := h.UserRepo.SetUser(user)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		response := resultDito.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(res).Encode(response)
	}

	res.WriteHeader(http.StatusOK)

	response := resultDito.SuccessResult{Code: http.StatusOK, Data: convertResponse(data)}
	json.NewEncoder(res).Encode(response)
	log.Println(data.ID)

}

func convertResponse(u models.User) usersdito.UserResponse {
	return usersdito.UserResponse{
		ID:    u.ID,
		Name:  u.Name,
		Email: u.Email,
		// Password: u.Password,
	}
}
