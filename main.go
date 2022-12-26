package main

import (
	"baharsah/helper/mysql"
	"baharsah/migration"
	"baharsah/router"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"

	"github.com/gorilla/mux"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	mysql.DBInit()
	migration.RunMigration()

	r := mux.NewRouter()
	router.Router(r.PathPrefix("/api/v1").Subrouter())
	log.Println("Server Running!")

	srverr := http.ListenAndServe("localhost:"+os.Getenv("PORT"), r)
	log.Println(srverr)
}
