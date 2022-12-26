package router

import (
	"baharsah/handlers"
	"baharsah/helper/mysql"
	"baharsah/middleware"
	"baharsah/repo"
	"net/http"

	"github.com/gorilla/mux"
)

func TrxRouter(r *mux.Router) {
	trxrepo := repo.RepoTRX(mysql.DB)
	h := handlers.HandlerTransaction(trxrepo)

	r.HandleFunc("/transaction/", middleware.Auth(h.SetTransaction)).Methods(http.MethodPost)
	r.HandleFunc("/order/", middleware.Auth(h.GetTransactions)).Methods(http.MethodGet)
	r.HandleFunc("/transaction/{id:[0-9]+}", middleware.Auth(h.GetTransaction)).Methods(http.MethodGet)
}
