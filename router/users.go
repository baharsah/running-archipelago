package router

import (
	"baharsah/handlers"
	"baharsah/pkg/mysql"
	"baharsah/repo"
	"net/http"

	"github.com/gorilla/mux"
)

func UserRoute(r *mux.Router) {
	userRepo := repo.RepoUser(mysql.DB)
	h := handlers.HandlerUser(userRepo)
	r.HandleFunc("/users", h.GetUsers).Methods(http.MethodGet)

}