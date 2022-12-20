package main

import (
	"baharsah/router"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()
	router.Router()

	http.ListenAndServe("localhost:5000", r)
}
