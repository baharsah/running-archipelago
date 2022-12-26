package handlers

import (
	resultDito "baharsah/dito/result"
	tripdito "baharsah/dito/trip"
	"baharsah/models"
	"baharsah/repo"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
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
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	res.Header().Set("Content-Type", "application/json")

	// minta data yang di def dari  dto trip
	datactx := req.Context().Value("file")
	filename := datactx.([]string)
	var collector = filename
	// request := new(tripdito.TripRequest)

	logrus.Println("Ini Body", req.Body)

	// try to allocate addressable value
	ctryid, _ := strconv.Atoi(req.FormValue("country_id"))
	qty, _ := strconv.Atoi(req.FormValue("qty"))
	day, _ := strconv.Atoi(req.FormValue("day"))
	night, _ := strconv.Atoi(req.FormValue("night"))

	request := tripdito.TripRequest{
		DestinationName: req.FormValue("destination_name"),
		Title:           req.FormValue("title"),
		CountryID:       ctryid,
		Accomodation:    req.FormValue("accomodation"),
		Transportation:  req.FormValue("transportation"),
		Eatenary:        req.FormValue("eat"),
		Day:             day,
		Night:           night,
		Quota:           uint(qty),
		Description:     req.FormValue("description"),
		DateTrip:        req.FormValue("date_trip"),
	}

	// konversi tanggal dan momen dalam form

	t, err := time.Parse("2006-01-02T15:04:05", request.DateTrip)
	dayTime := day * 24
	nightTime := night * 12
	toDateParse := t.Add(time.Duration(dayTime) + time.Duration(nightTime))

	// ambil meta gambar kedalam database

	log.Println(request)

	result := make([]models.ImageTrips, len(collector))
	for i, v := range collector {
		result[i] = models.ImageTrips{URL: os.Getenv("CDN_URL") + v}
		// logrus.Println("wow", v)
	}
	var ref []models.ImageTrips = result

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
		Quantity:        uint(qty),
		ImageTrips:      ref,
	}

	logrus.Println("ini hasil jadi", trip)

	tripdata, err := h.TripRepo.CreateTrip(trip)

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

func (h *tripRepoHandler) DeleteTrip(res http.ResponseWriter, req *http.Request) {

	idint, _ := strconv.Atoi(mux.Vars(req)["id"])
	trip, err := h.TripRepo.GetTrip(int(idint))

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		response := resultDito.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(res).Encode(response)
		return
	}

	data, errv := h.TripRepo.DeleteTrip(trip)

	if errv != nil {
		res.WriteHeader(http.StatusInternalServerError)
		response := resultDito.ErrorResult{Code: http.StatusInternalServerError, Message: errv.Error()}
		json.NewEncoder(res).Encode(response)
		return
	}

	res.WriteHeader(http.StatusOK)
	response := resultDito.SuccessResult{Code: http.StatusOK, Data: data}
	json.NewEncoder(res).Encode(response)

}
