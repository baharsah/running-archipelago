package router

import (
	"baharsah/handlers"
	"baharsah/helper/mysql"
	"baharsah/middleware"
	"baharsah/repo"
	"net/http"

	"github.com/gorilla/mux"
)

func TripRoute(r *mux.Router) {
	tripRepo := repo.RepoTrip(mysql.DB)
	h := handlers.HandlerTrip(tripRepo)

	r.HandleFunc("/trips", h.GetTrips).Methods(http.MethodGet)
	r.HandleFunc("/trip", middleware.UploadFilesTrip(h.SetTrip)).Methods(http.MethodPost)
	r.HandleFunc("/trip/{id}", h.DeleteTrip).Methods(http.MethodDelete)
}
