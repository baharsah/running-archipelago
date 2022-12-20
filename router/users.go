package router

import (
	"baharsah/handlers"

	"github.com/gorilla/mux"
)

func UserRoute(r *mux.Router) {
	r.HandleFunc("/users/:id", handlers.GetUsers)

}
