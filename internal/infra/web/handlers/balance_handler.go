package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/rafaelmgr12/go-pay/internal/usecase/transactions"
)

type WebBalanceHandler struct {
	TransactionUsecase transactions.TransactionsUseCase
}

func NewWebBalanceHandler(transactionUsecase transactions.TransactionsUseCase) *WebBalanceHandler {
	return &WebBalanceHandler{
		TransactionUsecase: transactionUsecase,
	}
}

func (h *WebBalanceHandler) GetBalanceHandle(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	userID := chi.URLParam(r, "_id")
	result, err := h.TransactionUsecase.BalanceGateway.GetAmountById(r.Context(), userID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error getting balance"))
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf(`{"balance": %f}`, result)))
}

func (h *WebBalanceHandler) TransferHandle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if !json.Valid(body) {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	var dto transactions.TransferInputDTO
	err = json.Unmarshal(body, &dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.TransactionUsecase.Transfer(dto.DebtorId, dto.CreditorId, dto.Amount, r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"message": "Transfer successful"}`))
}
