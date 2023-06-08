package handlers

import (
	"fmt"
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
