package main

import (
	"baharsah/helper/mysql"
	"baharsah/migration"
	"baharsah/router"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/gorilla/mux"
)

func main() {
	mysql.DBInit()
	migration.RunMigration()

	r := mux.NewRouter()
	router.Router(r.PathPrefix("/api/v1").Subrouter())
	log.Println("Server Running!")

	srverr := http.ListenAndServe("localhost:5001", r)
	log.Println(srverr)
}
