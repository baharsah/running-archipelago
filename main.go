package main

import (
	"baharsah/migration"
	"baharsah/pkg/mysql"
	"baharsah/router"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	mysql.DBInit()
	migration.RunMigration()

	r := mux.NewRouter()
	router.Router(r.PathPrefix("/api/v1").Subrouter())

	http.ListenAndServe("localhost:5000", r)
	log.Println("Server Running!")
}
