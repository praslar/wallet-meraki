package handler

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
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
			"error": "unauthorized role",
		})
		return
	}
	//// Extract the user ID from the token
	//secret := config.LoadEnv().Secret
	//tokenparse, err := jwt.Parse(token[1], func(token *jwt.Token) (interface{}, error) {
	//	// Use the same secret key that was used to sign the token
	//	return []byte(secret), nil
	//})
	//if err != nil || !tokenparse.Valid {
	//	http.Error(w, "Invalid Authorization header", http.StatusUnauthorized)
	//	return
	//}
	//
	//claims, ok := tokenparse.Claims.(jwt.MapClaims)
	//if !ok {
	//	http.Error(w, "Invalid token claims", http.StatusInternalServerError)
	//	return
	//}
	//XuserID, ok := claims["x-user-id"].(uuid.UUID)
	//if !ok {
	//	http.Error(w, "User ID not found in token claims", http.StatusInternalServerError)
	//	return
	//}
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
func convertToUUID(userIDStr string) (uuid.UUID, error) {
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		return uuid.Nil, err
	}
	return userID, nil
}