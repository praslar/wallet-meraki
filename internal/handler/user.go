package handler

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"net/http"
	"strings"
	"wallet/internal/model"
	"wallet/internal/service"
)

type UserHandler struct {
	userService   service.UserService
	authService   service.AuthService
	walletService service.WalletService
	tokenService  service.TokenService
}

func NewUserHandler(userService service.UserService, authService service.AuthService, walletService service.WalletService, tokenService service.TokenService) UserHandler {
	return UserHandler{
		userService:   userService,
		authService:   authService,
		walletService: walletService,
		tokenService:  tokenService,
	}
}

func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	requestUser := model.UserRequest{}
	w.Header().Set("Content-Type", "application/json")

	err := json.NewDecoder(r.Body).Decode(&requestUser)
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

	if err := h.userService.Register(requestUser.Email, requestUser.Password); err != nil {
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

	if err = json.NewEncoder(w).Encode(requestUser); err != nil {
		return
	}
}

func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	requestUser := model.UserRequest{}
	err := json.NewDecoder(r.Body).Decode(&requestUser)
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

	token, err := h.userService.Login(requestUser.Email, requestUser.Password)
	if err != nil {
		logrus.Errorf("Failed to login: %v", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		err := json.NewEncoder(w).Encode(map[string]interface{}{
			"error": "failed to login",
		})
		if err != nil {
			return
		}
		return
	}

	if err = json.NewEncoder(w).Encode(map[string]interface{}{
		"token": token,
	}); err != nil {
		return
	}
}

func (h *UserHandler) GetAllUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	jwtToken := r.Header.Get("Authorization")
	token := strings.Split(jwtToken, " ")

	if token[0] != "Bearer" {
		w.WriteHeader(http.StatusUnauthorized)
		err := json.NewEncoder(w).Encode(map[string]interface{}{
			"error": "unauthorized",
		})
		if err != nil {
			return
		}
		return
	}

	//jwtToken
	//tokendash := strings.Split(token[1], ".")
	if err := h.authService.ValidJWTToken(token[1], "admin"); err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		err := json.NewEncoder(w).Encode(map[string]interface{}{
			"error": "fail authorized",
		})
		if err != nil {
			return
		}
		return
	}

	users, err := h.userService.GetAllUser()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		err := json.NewEncoder(w).Encode(map[string]interface{}{
			"error": err.Error(),
		})
		if err != nil {
			return
		}
		return
	}
	if err = json.NewEncoder(w).Encode(map[string]interface{}{
		"data": users,
	}); err != nil {
		return
	}
}

func (h *UserHandler) CreateWallet(w http.ResponseWriter, r *http.Request) {
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

	//jwtToken := r.Header.Get("Authorization")
	//token := strings.Split(jwtToken, " ")
	//if token[0] != "Bearer" {
	//	w.WriteHeader(http.StatusUnauthorized)
	//	err := json.NewEncoder(w).Encode(map[string]interface{}{
	//		"error": "unauthorized jwt",
	//	})
	//	if err != nil {
	//		return
	//	}
	//	return
	//}
	//
	//// jwtToken
	//if err := h.authService.ValidJWTToken(token[1], "admin"); err != nil {
	//	w.WriteHeader(http.StatusUnauthorized)
	//	err := json.NewEncoder(w).Encode(map[string]interface{}{
	//		"error": "unauthorized valid",
	//	})
	//	if err != nil {
	//		return
	//	}
	//	return
	//}

	if err := h.walletService.CreateWallet(requestWallet.Name, requestWallet.UserID); err != nil {
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

	if err = json.NewEncoder(w).Encode(requestWallet); err != nil {
		return
	}
}

//
//func (h *UserHandler) CreateTokenAd(w http.ResponseWriter, r *http.Request) {
//	requestToken := model.TokenRequest{}
//	w.Header().Set("Content-Type", "application/json")
//
//	err := json.NewDecoder(r.Body).Decode(&requestToken)
//	if err != nil {
//		logrus.Errorf("Failed to get request body: %v", err.Error())
//		w.WriteHeader(http.StatusBadRequest)
//		err := json.NewEncoder(w).Encode(map[string]interface{}{
//			"error": err.Error(),
//		})
//		if err != nil {
//			return
//		}
//		return
//	}
//
//	jwtToken := r.Header.Get("Authorization")
//	token := strings.Split(jwtToken, " ")
//	if token[0] != "Bearer" {
//		w.WriteHeader(http.StatusUnauthorized)
//		err := json.NewEncoder(w).Encode(map[string]interface{}{
//			"error": "unauthorized jwt",
//		})
//		if err != nil {
//			return
//		}
//		return
//	}
//
//	// jwtToken
//	if err := h.authService.ValidJWTToken(token[1], "admin"); err != nil {
//		w.WriteHeader(http.StatusUnauthorized)
//		err := json.NewEncoder(w).Encode(map[string]interface{}{
//			"error": "unauthorized valid",
//		})
//		if err != nil {
//			return
//		}
//		return
//	}
//
//	if err := h.tokenService.CreateTokenAd(requestToken.WalletAddress, requestToken.Symbol, requestToken.Amount); err != nil {
//		logrus.Errorf("Failed create user: %v", err.Error())
//		w.WriteHeader(http.StatusInternalServerError)
//		err := json.NewEncoder(w).Encode(map[string]interface{}{
//			"error": err.Error(),
//		})
//		if err != nil {
//			return
//		}
//		return
//	}
//
//	if err = json.NewEncoder(w).Encode(requestToken); err != nil {
//		return
//	}
//}
//
//func (h *UserHandler) UpdateTokenAd(w http.ResponseWriter, r *http.Request) {
//	requestToken := model.TokenRequest{}
//	w.Header().Set("Content-Type", "application/json")
//
//	err := json.NewDecoder(r.Body).Decode(&requestToken)
//	if err != nil {
//		logrus.Errorf("Failed to get request body: %v", err.Error())
//		w.WriteHeader(http.StatusBadRequest)
//		err := json.NewEncoder(w).Encode(map[string]interface{}{
//			"error": err.Error(),
//		})
//		if err != nil {
//			return
//		}
//		return
//	}
//
//	jwtToken := r.Header.Get("Authorization")
//	token := strings.Split(jwtToken, " ")
//	if token[0] != "Bearer" {
//		w.WriteHeader(http.StatusUnauthorized)
//		err := json.NewEncoder(w).Encode(map[string]interface{}{
//			"error": "unauthorized jwt",
//		})
//		if err != nil {
//			return
//		}
//		return
//	}
//
//	// jwtToken
//	if err := h.authService.ValidJWTToken(token[1], "admin"); err != nil {
//		w.WriteHeader(http.StatusUnauthorized)
//		err := json.NewEncoder(w).Encode(map[string]interface{}{
//			"error": "unauthorized valid",
//		})
//		if err != nil {
//			return
//		}
//		return
//	}
//
//	if err := h.tokenService.UpdateTokenAd(requestToken.WalletAddress, requestToken.TokenID, requestToken.Symbol); err != nil {
//		logrus.Errorf("Failed create user: %v", err.Error())
//		w.WriteHeader(http.StatusInternalServerError)
//		err := json.NewEncoder(w).Encode(map[string]interface{}{
//			"error": err.Error(),
//		})
//		if err != nil {
//			return
//		}
//		return
//	}
//
//	if err = json.NewEncoder(w).Encode(requestToken); err != nil {
//		return
//	}
//}
//
//func (h *UserHandler) DeleteTokenAd(w http.ResponseWriter, r *http.Request) {
//	requestToken := model.TokenRequest{}
//	w.Header().Set("Content-Type", "application/json")
//
//	err := json.NewDecoder(r.Body).Decode(&requestToken)
//	if err != nil {
//		logrus.Errorf("Failed to get request body: %v", err.Error())
//		w.WriteHeader(http.StatusBadRequest)
//		err := json.NewEncoder(w).Encode(map[string]interface{}{
//			"error": err.Error(),
//		})
//		if err != nil {
//			return
//		}
//		return
//	}
//
//	jwtToken := r.Header.Get("Authorization")
//	token := strings.Split(jwtToken, " ")
//	if token[0] != "Bearer" {
//		w.WriteHeader(http.StatusUnauthorized)
//		err := json.NewEncoder(w).Encode(map[string]interface{}{
//			"error": "unauthorized jwt",
//		})
//		if err != nil {
//			return
//		}
//		return
//	}
//
//	// jwtToken
//	if err := h.authService.ValidJWTToken(token[1], "admin"); err != nil {
//		w.WriteHeader(http.StatusUnauthorized)
//		err := json.NewEncoder(w).Encode(map[string]interface{}{
//			"error": "unauthorized valid",
//		})
//		if err != nil {
//			return
//		}
//		return
//	}
//
//	if err := h.tokenService.DeleteTokenAd(requestToken.WalletAddress, requestToken.TokenID); err != nil {
//		logrus.Errorf("Failed create user: %v", err.Error())
//		w.WriteHeader(http.StatusInternalServerError)
//		err := json.NewEncoder(w).Encode(map[string]interface{}{
//			"error": err.Error(),
//		})
//		if err != nil {
//			return
//		}
//		return
//	}
//
//	if err = json.NewEncoder(w).Encode(requestToken); err != nil {
//		return
//	}
//}
//
//func (h *UserHandler) TransferTokenAd(w http.ResponseWriter, r *http.Request) {
//	requestToken := model.TokenRequest{}
//	requestTransaction := model.TransactionRequest{}
//
//	w.Header().Set("Content-Type", "application/json")
//
//	err := json.NewDecoder(r.Body).Decode(&requestToken)
//	if err != nil {
//		logrus.Errorf("Failed to get request body: %v", err.Error())
//		w.WriteHeader(http.StatusBadRequest)
//		err := json.NewEncoder(w).Encode(map[string]interface{}{
//			"error": err.Error(),
//		})
//		if err != nil {
//			return
//		}
//		return
//	}
//
//	jwtToken := r.Header.Get("Authorization")
//	token := strings.Split(jwtToken, " ")
//	if token[0] != "Bearer" {
//		w.WriteHeader(http.StatusUnauthorized)
//		err := json.NewEncoder(w).Encode(map[string]interface{}{
//			"error": "unauthorized jwt",
//		})
//		if err != nil {
//			return
//		}
//		return
//	}
//
//	// jwtToken
//	if err := h.authService.ValidJWTToken(token[1], "admin"); err != nil {
//		w.WriteHeader(http.StatusUnauthorized)
//		err := json.NewEncoder(w).Encode(map[string]interface{}{
//			"error": "unauthorized valid",
//		})
//		if err != nil {
//			return
//		}
//		return
//	}
//
//	if err := h.tokenService.TransferTokenAd(requestTransaction.SenderWalletAddress, requestTransaction.ReceiverWalletAddress, requestToken.TokenID, requestToken.Amount); err != nil {
//		logrus.Errorf("Failed create user: %v", err.Error())
//		w.WriteHeader(http.StatusInternalServerError)
//		err := json.NewEncoder(w).Encode(map[string]interface{}{
//			"error": err.Error(),
//		})
//		if err != nil {
//			return
//		}
//		return
//	}
//
//	if err = json.NewEncoder(w).Encode(requestToken); err != nil {
//		return
//	}
//}
