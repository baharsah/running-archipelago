package handlers

import (
	authDito "baharsah/dito/auth"
	resultDito "baharsah/dito/result"
	usersdito "baharsah/dito/users"
	"baharsah/helper/bcrypt"
	"baharsah/helper/jawatoken"
	"baharsah/models"
	"baharsah/repo"
	userValdilator "baharsah/valdilator/uservaldilator"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
)

type userRepoHandler struct {
	UserRepo repo.UserRepo
}

func HandlerUser(UserRepo repo.UserRepo) *userRepoHandler {
	return &userRepoHandler{UserRepo}

}

func (h *userRepoHandler) GetUsers(w http.ResponseWriter, r *http.Request) {

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

func (h *userRepoHandler) CreateUser(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	request := new(authDito.RegisterRequest)

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
		Address:  request.Address,
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

func (h *userRepoHandler) AuthUser(res http.ResponseWriter, req *http.Request) {

	res.Header().Set("Content-Type", "application/json")
	request := new(authDito.LoginRequest)
	if err := json.NewDecoder(req.Body).Decode(&request); err != nil {
		res.WriteHeader(http.StatusBadRequest)
		response := resultDito.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(res).Encode(response)
		return
	}

	user := models.User{
		Email:    request.Email,
		Password: request.Password,
	}

	user, err := h.UserRepo.GetUser(user)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		response := resultDito.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(res).Encode(response)
		return
	}
	// log.Println("request", request.Password)
	// log.Println("user", user.Password)

	isValidBcrypt := bcrypt.CheckPasswordHash(request.Password, user.Password)

	if !isValidBcrypt {
		res.WriteHeader(http.StatusUnauthorized)
		response := resultDito.ErrorResult{Code: http.StatusUnauthorized, Message: "Invalid password"}
		json.NewEncoder(res).Encode(response)
		return
	}

	claims := jwt.MapClaims{}
	claims["email"] = user.Email
	claims["isAdmin"] = user.IsAdmin
	claims["userID"] = user.ID
	claims["exp"] = time.Now().Add(time.Hour * 2).Unix() // 2 jam expired

	token, errGenerateToken := jawatoken.GenerateToken(&claims)
	if errGenerateToken != nil {
		log.Println(errGenerateToken)
		log.Println("Unauthorize")
		return
	}

	loginResponse := authDito.LoginResponse{
		Email:    user.Email,
		Password: user.Password,
		Token:    token,
		IsAdmin:  user.IsAdmin,
	}

	res.Header().Set("Content-Type", "application/json")
	response := resultDito.SuccessResult{Code: http.StatusOK, Data: loginResponse}
	json.NewEncoder(res).Encode(response)

}

func convertResponse(u models.User) usersdito.UserResponse {
	return usersdito.UserResponse{
		ID:    u.ID,
		Name:  u.Name,
		Email: u.Email,
		// Password: u.Password,
	}
}

func (h *userRepoHandler) GetUser(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	id, _ := strconv.Atoi(mux.Vars(req)["id"])
	log.Println(id)
	user, err := h.UserRepo.GetUserID(models.User{ID: id})
	// h.UserRepo.GetUser(models.User{ID: })
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		response := resultDito.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(res).Encode(response)
		return
	}
	response := resultDito.SuccessResult{Code: http.StatusOK, Data: user}
	json.NewEncoder(res).Encode(response)
}
