package handlers

import (
	dito "baharsah/dito/result"
	"baharsah/repo"
	"encoding/json"
	"net/http"
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
	response := dito.SuccessResult{Code: http.StatusOK, Data: users}
	json.NewEncoder(w).Encode(response)
}
