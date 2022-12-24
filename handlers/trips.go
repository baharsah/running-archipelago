package handlers

import (
	resultDito "baharsah/dito/result"
	tripdito "baharsah/dito/trip"
	"baharsah/models"
	"baharsah/repo"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
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

func (h *tripRepoHandler) SetTrip(res http.ResponseWriter, req *http.Request) {
	// minta data yang di def dari  dto trip

	request := new(tripdito.TripRequest)

	if err := json.NewDecoder(req.Body).Decode(&request); err != nil {

		res.WriteHeader(http.StatusBadRequest)
		response := resultDito.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(res).Encode(response)
		return

	}

	validation := validator.New()

	err := validation.Struct(request)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		response := resultDito.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(res).Encode(response)
		return
	}

	// konversi tanggal dan momen dalam form

	t, err := time.Parse("2006-01-02T15:04:05", request.DateTrip)
	dayTime := request.Day * 24
	nightTime := request.Night * 12
	toDateParse := t.Add(time.Duration(dayTime) + time.Duration(nightTime))

	// ambil meta gambar kedalam database
	datactx := req.Context().Value("file")
	filename := datactx.([]models.ImageTrips)

	// rubah data array menjadi struct

	trip := models.Trips{

		DestinationName: request.DestinationName,
		Title:           request.Title,
		CtryID:          uint(request.CountryID),
		Accomodation:    request.Accomodation,
		Transportation:  request.Transportation,
		Eatenary:        request.Eatenary,
		FromDate:        t,
		ToDate:          toDateParse,
		Description:     request.Description,
		Quantity:        request.Quota,
		ImageTrips:      filename,
	}
	logrus.Println(trip)

	tripdata, err := h.TripRepo.CreateTrip(trip)
	log.Println("ini hasil result", tripdata)

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		response := resultDito.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(res).Encode(response)
		return
	}

	res.WriteHeader(http.StatusOK)
	response := resultDito.SuccessResult{Code: http.StatusOK, Data: tripdata}
	json.NewEncoder(res).Encode(response)

}
