package handler

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"net/http"
	"wallet/internal/model"
	"wallet/internal/service"
)

type TokenHandler struct {
	tokenService service.TokenServiceInterface
}

func NewTokenHandler(tokenService service.TokenServiceInterface) TokenHandler {
	return TokenHandler{
		tokenService: tokenService,
	}
}

func (h *TokenHandler) CreateToken(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	requestToken := model.TokenRequest{}
	err := json.NewDecoder(r.Body).Decode(&requestToken)
	if err != nil {
		logrus.Errorf("Failed to get request body: %v", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"error": err.Error(),
		})
		return
	}

	if err := h.tokenService.CreateToken(requestToken.Symbol, requestToken.Price); err != nil {
		logrus.Errorf("Failed create token: %v", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		err := json.NewEncoder(w).Encode(map[string]interface{}{
			"error": err.Error(),
		})
		if err != nil {
			return
		}
		return
	}
	if err = json.NewEncoder(w).Encode(requestToken); err != nil {

		logrus.Errorf("Failed to get request body: %v", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"error": err.Error(),
		})
		return
	}
}

func (h *TokenHandler) DeleteToken(w http.ResponseWriter, r *http.Request) {
	requestToken := model.TokenRequest{}
	w.Header().Set("Content-Type", "application/json")

	err := json.NewDecoder(r.Body).Decode(&requestToken)
	if err != nil {
		logrus.Errorf("Failed to get request body: %v", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		err := json.NewEncoder(w).Encode(map[string]interface{}{
			"error": err.Error(),
		})
		if err != nil {
			return
		}
		return
	}

	if err := h.tokenService.DeleteToken(requestToken.Address); err != nil {
		logrus.Errorf("Failed delete token: %v", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		err := json.NewEncoder(w).Encode(map[string]interface{}{
			"error": err.Error(),
		})
		if err != nil {
			return
		}
		return
	}

	if err = json.NewEncoder(w).Encode(requestToken); err != nil {
		logrus.Errorf("Failed to get request body: %v", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"error": err.Error(),
		})
		return
	}
}

func (h *TokenHandler) UpdateToken(w http.ResponseWriter, r *http.Request) {
	requestToken := model.TokenRequest{}
	w.Header().Set("Content-Type", "application/json")

	err := json.NewDecoder(r.Body).Decode(&requestToken)
	if err != nil {
		logrus.Errorf("Failed to get request body: %v", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		err := json.NewEncoder(w).Encode(map[string]interface{}{
			"error": err.Error(),
		})
		if err != nil {
			return
		}
		return
	}

	if err := h.tokenService.UpdateToken(requestToken.Address, requestToken.Symbol, requestToken.Price); err != nil {
		logrus.Errorf("Failed create user: %v", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		err := json.NewEncoder(w).Encode(map[string]interface{}{
			"error": err.Error(),
		})
		if err != nil {
			return
		}
		return
	}

	if err = json.NewEncoder(w).Encode(requestToken); err != nil {

		logrus.Errorf("Failed to get request body: %v", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"error": err.Error(),
		})
		return
	}
}

func (h *TokenHandler) SendUserToken(w http.ResponseWriter, r *http.Request) {
	requestTransaction := model.TransactionRequest{}

	w.Header().Set("Content-Type", "application/json")

	err := json.NewDecoder(r.Body).Decode(&requestTransaction)
	if err != nil {
		logrus.Errorf("Failed to get request body: %v", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		err := json.NewEncoder(w).Encode(map[string]interface{}{
			"error": err.Error(),
		})
		if err != nil {
			return
		}
		return
	}

	//if err := h.tokenService.SendUserToken(requestTransaction.SenderWalletAddress, requestTransaction.ReceiverWalletAddress, requestTransaction.TokenAddress, requestTransaction.Amount); err != nil {
	//	logrus.Errorf("Failed create user: %v", err.Error())
	//	w.WriteHeader(http.StatusInternalServerError)
	//	err := json.NewEncoder(w).Encode(map[string]interface{}{
	//		"error": err.Error(),
	//	})
	//	if err != nil {
	//		return
	//	}
	//	return
	//}

	if err = json.NewEncoder(w).Encode(requestTransaction); err != nil {

		logrus.Errorf("Failed to get request body: %v", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"error": err.Error(),
		})

		return
	}

}
