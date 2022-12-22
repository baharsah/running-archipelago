package router

import (
	"baharsah/handlers"
	"baharsah/helper/mysql"
	"baharsah/repo"
	"net/http"

	"github.com/gorilla/mux"
)

func UserRoute(r *mux.Router) {
	userRepo := repo.RepoUser(mysql.DB)
	h := handlers.HandlerUser(userRepo)
	r.HandleFunc("/users", h.GetUsers).Methods(http.MethodGet)
	r.HandleFunc("/register", h.CreateUser).Methods(http.MethodPost)

}
