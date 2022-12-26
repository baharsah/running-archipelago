package router

import (
	"baharsah/handlers"
	"baharsah/helper/mysql"
	"baharsah/middleware"
	"baharsah/repo"
	"net/http"

	"github.com/gorilla/mux"
)

func CtryRouter(r *mux.Router) {
	ctryRepo := repo.RepoCountry(mysql.DB)
	h := handlers.HandlerCountry(ctryRepo)

	r.HandleFunc("/country", middleware.Auth(h.SetCountry)).Methods(http.MethodPost)
	r.HandleFunc("/country", middleware.Auth(h.GetCountries)).Methods(http.MethodGet)
	r.HandleFunc("/country/{id}", middleware.Auth(h.GetCountry)).Methods(http.MethodGet)
	r.HandleFunc("/country/{id}", middleware.Auth(h.DeleteCountry)).Methods(http.MethodDelete)
	r.HandleFunc("/country/{id}", middleware.Auth(h.UpdateCountry)).Methods(http.MethodPatch)
}
