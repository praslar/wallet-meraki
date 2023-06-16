package handler

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"net/http"
	"wallet/internal/model"
	"wallet/internal/service"
)

type WalletHandler struct {
	WalletService service.WalletService
}

func NewWalletHandler(WalletService service.WalletService) WalletHandler {
	return WalletHandler{
		WalletService: WalletService,
	}
}

func (h *WalletHandler) CreateWallet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	requestWallet := model.WalletRequest{}
	w.Header().Set("Content-Type", "application/json")

	err := json.NewDecoder(r.Body).Decode(&requestWallet)
	if err != nil {
		logrus.Errorf("Failed to get request body: %v", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"error": err.Error(),
		})
		return
	}

	//currentUser := r.Header.Get("x-user-id")
	//nameWallet := r.Header.Get("name")
	if err := h.WalletService.CreateWallet(requestWallet.UserID, requestWallet.Name); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logrus.Errorf(err.Error())
		json.NewEncoder(w).Encode(map[string]interface{}{
			"error": "Internal server error",
		})
		return
	}

}
