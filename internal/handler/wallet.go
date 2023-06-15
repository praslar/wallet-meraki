package handler

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"net/http"
	"strings"
	"wallet/internal/model"
	"wallet/internal/service"
)

type WalletHandler struct {
	WalletService service.WalletService
	authService   service.AuthService
}

func NewWalletHandler(WalletService service.WalletService, authService service.AuthService) WalletHandler {
	return WalletHandler{
		WalletService: WalletService,
		authService:   authService,
	}
}

func (h *WalletHandler) CreateWallet(w http.ResponseWriter, r *http.Request) {
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

	jwtToken := r.Header.Get("Authorization")
	token := strings.Split(jwtToken, " ")
	if token[0] != "Bearer" {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"error": "unauthorized jwt",
		})
		return
	}

	// jwtToken
	if err := h.authService.ValidJWTToken(token[1], "user"); err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"error": "unauthorized role",
		})
		return
	}

	if err := h.authService.ValidJWTTokenXuser(token[1], requestWallet.UserID); err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"error": "unauthorized role",
		})
		return
	}

	if err := h.WalletService.CreateWallet(requestWallet.Name, requestWallet.UserID); err != nil {
		logrus.Errorf("Failed create wallet: %v", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"error": err.Error(),
		})
		return
	}

	if err = json.NewEncoder(w).Encode(requestWallet); err != nil {
		return
	}
}
func (h *WalletHandler) GetOneWallet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	requestWallet := model.WalletRequest{}
	err := json.NewDecoder(r.Body).Decode(&requestWallet)
	if err != nil {
		logrus.Errorf("Failed to get request body: %v", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"error": err.Error(),
		})
		return
	}
	//name := r.URL.Query().Get("name")
	//userIDstr := r.URL.Query().Get("user_id")

	jwtToken := r.Header.Get("Authorization")
	token := strings.Split(jwtToken, " ")
	if token[0] != "Bearer" {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"error": "unauthorized bearer",
		})
		return
	}

	// jwtToken
	if err := h.authService.ValidJWTToken(token[1], "user"); err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"error": "unauthorized jwt",
		})
		return
	}
	if err := h.authService.ValidJWTTokenXuser(token[1], requestWallet.UserID); err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"error": "unauthorized jwt",
		})
		return
	}

	rs, err := h.WalletService.GetOneWallet(requestWallet.Name, requestWallet.UserID)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"error": "unauthorized jwt",
		})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(rs)
}
