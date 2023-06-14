package handler

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
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
			"error": "unauthorized valid",
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

	name := r.URL.Query().Get("name")

	rs, err := h.WalletService.GetOneWallet(name)
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

func (h *WalletHandler) GetAllWallet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

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

	order := r.URL.Query().Get("order")
	name := r.URL.Query().Get("name")
	userID := r.URL.Query().Get("user_id")

	pageSize := r.URL.Query().Get("page_size")
	pageSizeInt, _ := strconv.Atoi(pageSize)

	page := r.URL.Query().Get("page")
	pageInt, _ := strconv.Atoi(page)

	wallets, err := h.WalletService.GetAllWallet(order, name, userID, pageSizeInt, pageInt)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"error": err.Error(),
		})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(wallets)
}

func (h *WalletHandler) DeleteWallet(w http.ResponseWriter, r *http.Request) {
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

	err = h.WalletService.DeleteWallet(requestWallet.Name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *WalletHandler) UpdateWallet(w http.ResponseWriter, r *http.Request) {
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

	wallet, err := h.WalletService.UpdateWallet(requestWallet.UserID, requestWallet.Name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "%v", wallet)
}
