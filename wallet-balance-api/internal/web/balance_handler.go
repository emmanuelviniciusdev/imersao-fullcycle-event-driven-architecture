package web

import (
	"encoding/json"
	"fmt"
	"github.com/emmanuelviniciusdev/imersao-fullcycle-event-driven-architecture/wallet-balance-api/internal/usecase/get_balance"
	"github.com/go-chi/chi/v5"
	"net/http"
)

type WebBalanceHandler struct {
	GetBalanceUsecase get_balance.GetBalanceUsecase
}

func NewWebBalanceHandler(getBalanceUsecase get_balance.GetBalanceUsecase) *WebBalanceHandler {
	return &WebBalanceHandler{GetBalanceUsecase: getBalanceUsecase}
}

func (h *WebBalanceHandler) GetBalance(w http.ResponseWriter, r *http.Request) {
	accountID := chi.URLParam(r, "account_id")

	output, err := h.GetBalanceUsecase.Execute(accountID)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(output)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
