package handlers

import (
	countriesdito "baharsah/dito/countries"
	resultDito "baharsah/dito/result"
	"baharsah/models"
	"baharsah/repo"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

type countriesRepoHandler struct {
	CountryRepo repo.CountryRepo
}

func HandlerCountry(CountryRepo repo.CountryRepo) *countriesRepoHandler {
	return &countriesRepoHandler{CountryRepo}
}

func (h *countriesRepoHandler) SetCountry(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	request := new(countriesdito.CountryRequest)

	if err := json.NewDecoder(req.Body).Decode(request); err != nil {
		response := resultDito.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(res).Encode(response)
		return
	}
	co := models.Countries{
		Country: request.Country,
	}
	country, err := h.CountryRepo.SetCountry(co)
	if err != nil {
		response := resultDito.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(res).Encode(response)
		return
	}
	resq := resultDito.SuccessResult{Code: http.StatusCreated, Data: countriesdito.CountryResponse{CountryID: int(country.IDCountries), Country: request.Country}}
	json.NewEncoder(res).Encode(resq)
}

func (h *countriesRepoHandler) GetCountries(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	ctry, err := h.CountryRepo.GetCounties()
	if err != nil {
		response := resultDito.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(res).Encode(response)
		return
	}
	response := resultDito.SuccessResult{Code: http.StatusOK, Data: ctry}
	json.NewEncoder(res).Encode(response)
}

func (h *countriesRepoHandler) GetCountry(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	id, _ := strconv.Atoi(mux.Vars(req)["id"])
	ctry, err := h.CountryRepo.GetCountry(id)
	if err != nil {
		response := resultDito.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(res).Encode(response)
		return
	}
	response := resultDito.SuccessResult{Code: http.StatusOK, Data: ctry}
	json.NewEncoder(res).Encode(response)
}

func (h *countriesRepoHandler) DeleteCountry(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	id, _ := strconv.Atoi(mux.Vars(req)["id"])
	_, err := h.CountryRepo.DeleteCountry(models.Countries{IDCountries: uint(id)})

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		response := resultDito.ErrorResult{Code: http.StatusInternalServerError, Message: strings.Split(err.Error(), ":")[0] + " : ID tersebut telah terhubung dengan trip. harap hapus terlebih dahulu trip yang berhubungan dengan negara ini."}
		json.NewEncoder(res).Encode(response)
		return

	}
	response := resultDito.SuccessResult{Code: http.StatusOK}
	json.NewEncoder(res).Encode(response)
}

func (h *countriesRepoHandler) UpdateCountry(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	id, _ := strconv.Atoi(mux.Vars(req)["id"])
	request := new(countriesdito.CountryRequest)
	if err := json.NewDecoder(req.Body).Decode(request); err != nil {
		response := resultDito.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(res).Encode(response)
		return
	}
	co := countriesdito.CountryRequest{
		Country:   request.Country,
		CountryID: id,
	}
	cor := models.Countries{
		Country:     co.Country,
		IDCountries: uint(co.CountryID),
	}

	country, err := h.CountryRepo.UpdateCountry(cor)
	if err != nil {
		response := resultDito.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(res).Encode(response)
		return
	}
	response := resultDito.SuccessResult{Code: http.StatusOK, Data: country}
	json.NewEncoder(res).Encode(response)
}
