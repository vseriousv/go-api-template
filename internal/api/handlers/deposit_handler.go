package handlers

import (
	"encoding/json"
	"github.com/vseriousv/go-api-template/internal/api/services"
	"net/http"
)

type DepositHandler struct {
	depositService *services.DepositService
}

func NewDepositHandler(depositService *services.DepositService) *DepositHandler {
	return &DepositHandler{
		depositService: depositService,
	}
}

func (h *DepositHandler) DepositCalculate(w http.ResponseWriter, r *http.Request) {
	data, err := h.depositService.GetDeposit()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
