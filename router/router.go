package router

import "github.com/gorilla/mux"

func Router(r *mux.Router) {

	UserRoute(r)

}