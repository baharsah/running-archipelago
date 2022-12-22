package main

import (
	"baharsah/helper/mysql"
	"baharsah/migration"
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

	http.ListenAndServe("localhost:5001", r)
	log.Println("Server Running!")
}
