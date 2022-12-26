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
	// I'll use this if needed because IDK
	r.HandleFunc("/users", h.GetUsers).Methods(http.MethodGet)
	r.HandleFunc("/users/{id}", h.GetUser).Methods(http.MethodGet)
	r.HandleFunc("/login", h.AuthUser).Methods(http.MethodPost)

	r.HandleFunc("/register", h.CreateUser).Methods(http.MethodPost)

}
