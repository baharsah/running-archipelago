package handlers

import (
	resultDito "baharsah/dito/result"
	transactionsdito "baharsah/dito/transactions"
	"baharsah/models"
	"baharsah/repo"
	"encoding/json"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"

	"net/http"
)

type transactionRepoHandler struct {
	TRXRepo repo.TransactionRepo
}

func HandlerTransaction(TRXRepo repo.TransactionRepo) *transactionRepoHandler {
	return &transactionRepoHandler{TRXRepo}
}

func (h *transactionRepoHandler) SetTransaction(res http.ResponseWriter, req *http.Request) {

	res.Header().Set("Content-Type", "application/json")

	request := new(transactionsdito.TransactionsRequest)
	if err := json.NewDecoder(req.Body).Decode(request); err != nil {
		res.WriteHeader(http.StatusBadRequest)
		response := resultDito.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(res).Encode(response)
		return
	}

	validation := validator.New()
	err := validation.Struct(request)

	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		response := resultDito.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(res).Encode(response)
		return
	}

	trxmodel := models.Transactions{
		UserID:        request.UserID,
		TripID:        request.TripID,
		PaymentStatus: 1,
		Quantity:      uint(request.Qty),
	}
	trx, err := h.TRXRepo.SetTransaction(trxmodel)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		response := resultDito.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(res).Encode(response)
		return
	}

	trxres, err := h.TRXRepo.GetTransaction(trx.ID)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		response := resultDito.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(res).Encode(response)
		return
	}

	res.WriteHeader(http.StatusCreated)
	// res.WriteHeader(http.StatusOK)
	response := resultDito.SuccessResult{Code: http.StatusOK, Data: trxres}
	json.NewEncoder(res).Encode(response)

}

func (h *transactionRepoHandler) GetTransactions(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	trxs, err := h.TRXRepo.GetTransactions()
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		response := resultDito.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(res).Encode(response)
	}

	response := resultDito.SuccessResult{Code: http.StatusOK, Data: trxs}
	json.NewEncoder(res).Encode(response)

}

func (h *transactionRepoHandler) GetTransaction(res http.ResponseWriter, req *http.Request) {

	res.Header().Set("Content-Type", "application/json")
	iddata, _ := strconv.Atoi(mux.Vars(req)["id"])

	trxdata, err := h.TRXRepo.GetTransaction(iddata)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		response := resultDito.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(res).Encode(response)
		return
	}
	response := resultDito.SuccessResult{Code: http.StatusOK, Data: trxdata}
	json.NewEncoder(res).Encode(response)

}
