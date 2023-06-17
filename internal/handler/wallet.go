package handler

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
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

//func (h *WalletHandler) CreateWallet(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Content-Type", "application/json")
//
//	requestWallet := model.WalletRequest{}
//	w.Header().Set("Content-Type", "application/json")
//
//	err := json.NewDecoder(r.Body).Decode(&requestWallet)
//	if err != nil {
//		logrus.Errorf("Failed to get request body: %v", err.Error())
//		w.WriteHeader(http.StatusBadRequest)
//		json.NewEncoder(w).Encode(map[string]interface{}{
//			"error": err.Error(),
//		})
//		return
//	}
//
//	if err := h.WalletService.CreateWallet(requestWallet.UserID, requestWallet.Name); err != nil {
//		w.WriteHeader(http.StatusInternalServerError)
//		logrus.Errorf(err.Error())
//		json.NewEncoder(w).Encode(map[string]interface{}{
//			"error": "Internal server error",
//		})
//		return
//	}
//
//}

func (h *WalletHandler) CreateWallet(w http.ResponseWriter, r *http.Request) {
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

	currentUser := r.Header.Get("x-user-id")
	if err := h.WalletService.CreateWallet(requestWallet.Name, currentUser); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logrus.Errorf(err.Error())
		json.NewEncoder(w).Encode(map[string]interface{}{
			"error": "Internal server error",
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Wallet created successfully",
	})
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
	currentUser := r.Header.Get("x-user-id")

	rs, err := h.WalletService.GetOneWallet(currentUser, requestWallet.Name)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"error": "unauthorized jwt",
		})
		return
	}

	if err := json.NewEncoder(w).Encode(map[string]interface{}{
		"data": rs,
	}); err != nil {
		return
	}
}

func (h *WalletHandler) GetAllWallet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

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
	if err = json.NewEncoder(w).Encode(map[string]interface{}{
		"data": wallets,
	}); err != nil {
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

	currentUser := r.Header.Get("x-user-id")
	wallet, err := h.WalletService.UpdateWallet(currentUser, requestWallet.Name, requestWallet.UpdateName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err = json.NewEncoder(w).Encode(map[string]interface{}{
		"data": wallet,
	}); err != nil {
		return
	}
}
