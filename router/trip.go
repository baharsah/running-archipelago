package router

import (
	"baharsah/handlers"
	"baharsah/helper/mysql"
	"baharsah/repo"
	"net/http"

	"github.com/gorilla/mux"
)

func TripRoute(r *mux.Router) {
	tripRepo := repo.RepoTrip(mysql.DB)
	h := handlers.HandlerTrip(tripRepo)

	r.HandleFunc("/trips", h.GetTrips).Methods(http.MethodGet)
}
