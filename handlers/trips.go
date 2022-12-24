package handlers

import (
	resultDito "baharsah/dito/result"
	"baharsah/repo"
	"encoding/json"
	"net/http"
)

type tripRepoHandler struct {
	TripRepo repo.TripRepo
}

func HandlerTrip(TripRepo repo.TripRepo) *tripRepoHandler {

	return &tripRepoHandler{TripRepo}
}

func (h *tripRepoHandler) GetTrips(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	trip, err := h.TripRepo.GetTrips()
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(err.Error())
	}
	response := resultDito.SuccessResult{Code: http.StatusOK, Data: trip}
	json.NewEncoder(res).Encode(response)

}
