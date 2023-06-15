package handler

import (
	"encoding/json"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"strings"
	"wallet/config"
	"wallet/internal/model"
	"wallet/internal/service"
	"wallet/internal/utils"
)

type UserHandler struct {
	userService      service.UserService
	authService      service.AuthService
	walletService    service.WalletService
	tokenService     service.TokenService
	coingeckoService service.CoingeckoService
}

func NewUserHandler(userService service.UserService, authService service.AuthService, walletService service.WalletService, tokenService service.TokenService, coingeckoService service.CoingeckoService) UserHandler {
	return UserHandler{
		userService:      userService,
		authService:      authService,
		walletService:    walletService,
		tokenService:     tokenService,
		coingeckoService: coingeckoService,
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
	if err := h.authService.ValidJWTToken(token[1], utils.Admin); err != nil {
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

// TODO : lấy x-userid từ jwt để khi tạo wallet nhập thẳng userid = x userid
func (h *UserHandler) CreateWallet(w http.ResponseWriter, r *http.Request) {
	secret := config.LoadEnv().Secret
	requestWallet := model.WalletRequest{}
	w.Header().Set("Content-Type", "application/json")

	err := json.NewDecoder(r.Body).Decode(&requestWallet)
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

	jwtToken := r.Header.Get("Authorization")
	token := strings.Split(jwtToken, " ")

	if token[0] != "Bearer" {
		w.WriteHeader(http.StatusUnauthorized)
		err := json.NewEncoder(w).Encode(map[string]interface{}{
			"error": "unauthorized jwt",
		})
		if err != nil {
			return
		}
		return
	}

	// jwtToken
	if err := h.authService.ValidJWTToken(token[1], utils.Admin); err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		err := json.NewEncoder(w).Encode(map[string]interface{}{
			"error": "unauthorized valid",
		})
		if err != nil {
			return
		}
		return
	}

	// Extract the user ID from the token

	tokenparse, err := jwt.Parse(token[1], func(token *jwt.Token) (interface{}, error) {
		// Use the same secret key that was used to sign the token
		return []byte(secret), nil
	})
	if err != nil || !tokenparse.Valid {
		http.Error(w, "Invalid Authorization header", http.StatusUnauthorized)
		return
	}

	claims, ok := tokenparse.Claims.(jwt.MapClaims)
	if !ok {
		http.Error(w, "Invalid token claims", http.StatusInternalServerError)
		return
	}
	XuserID, ok := claims["x-user-id"].(uuid.UUID)
	if !ok {
		http.Error(w, "User ID not found in token claims", http.StatusInternalServerError)
		return
	}

	if err := h.walletService.CreateWallet(requestWallet.Name, XuserID); err != nil {
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

func (h *UserHandler) CreateToken(w http.ResponseWriter, r *http.Request) {
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
	jwtToken := r.Header.Get("Authorization")
	token := strings.Split(jwtToken, " ")
	if token[0] != "Bearer" {
		w.WriteHeader(http.StatusUnauthorized)
		err := json.NewEncoder(w).Encode(map[string]interface{}{
			"error": "unauthorized jwt",
		})
		if err != nil {
			return
		}
		return
	}
	// jwtToken
	if err := h.authService.ValidJWTToken(token[1], utils.Admin); err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		err := json.NewEncoder(w).Encode(map[string]interface{}{
			"error": "unauthorized valid",
		})
		if err != nil {
			return
		}
		return
	}
	// Extract the user ID from the token

	if err := h.tokenService.CreateToken(requestToken.Symbol, requestToken.Price); err != nil {
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
		return
	}
}

func (h *UserHandler) DeleteToken(w http.ResponseWriter, r *http.Request) {
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

	jwtToken := r.Header.Get("Authorization")
	token := strings.Split(jwtToken, " ")
	if token[0] != "Bearer" {
		w.WriteHeader(http.StatusUnauthorized)
		err := json.NewEncoder(w).Encode(map[string]interface{}{
			"error": "unauthorized jwt",
		})
		if err != nil {
			return
		}
		return
	}

	// jwtToken
	if err := h.authService.ValidJWTToken(token[1], utils.Admin); err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		err := json.NewEncoder(w).Encode(map[string]interface{}{
			"error": "unauthorized valid",
		})
		if err != nil {
			return
		}
		return
	}

	if err := h.tokenService.DeleteToken(requestToken.Address); err != nil {
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
		return
	}
}

func (h *UserHandler) UpdateToken(w http.ResponseWriter, r *http.Request) {
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

	jwtToken := r.Header.Get("Authorization")
	token := strings.Split(jwtToken, " ")
	if token[0] != "Bearer" {
		w.WriteHeader(http.StatusUnauthorized)
		err := json.NewEncoder(w).Encode(map[string]interface{}{
			"error": "unauthorized jwt",
		})
		if err != nil {
			return
		}
		return
	}

	// jwtToken
	if err := h.authService.ValidJWTToken(token[1], utils.Admin); err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		err := json.NewEncoder(w).Encode(map[string]interface{}{
			"error": "unauthorized valid",
		})
		if err != nil {
			return
		}
		return
	}

	if err := h.tokenService.UpdateToken(requestToken.Address); err != nil {
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
		return
	}
}

func (h *UserHandler) SendUserToken(w http.ResponseWriter, r *http.Request) {
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

	jwtToken := r.Header.Get("Authorization")
	token := strings.Split(jwtToken, " ")
	if token[0] != "Bearer" {
		w.WriteHeader(http.StatusUnauthorized)
		err := json.NewEncoder(w).Encode(map[string]interface{}{
			"error": "unauthorized jwt",
		})
		if err != nil {
			return
		}
		return
	}

	// jwtToken
	if err := h.authService.ValidJWTToken(token[1], utils.Admin); err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		err := json.NewEncoder(w).Encode(map[string]interface{}{
			"error": "unauthorized valid",
		})
		if err != nil {
			return
		}
		return
	}

	if err := h.tokenService.SendUserToken(requestTransaction.SenderWalletAddress, requestTransaction.ReceiverWalletAddress, requestTransaction.TokenAddress, requestTransaction.Amount); err != nil {
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

	if err = json.NewEncoder(w).Encode(requestTransaction); err != nil {
		return
	}
}

func (h *UserHandler) GetCoinInfo(w http.ResponseWriter, r *http.Request) {

	jwtToken := r.Header.Get("Authorization")
	token := strings.Split(jwtToken, " ")
	if token[0] != "Bearer" {
		w.WriteHeader(http.StatusUnauthorized)
		err := json.NewEncoder(w).Encode(map[string]interface{}{
			"error": "unauthorized jwt",
		})
		if err != nil {
			return
		}
		return
	}

	// jwtToken
	if err := h.authService.ValidJWTToken(token[1], utils.Admin); err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		err := json.NewEncoder(w).Encode(map[string]interface{}{
			"error": "unauthorized valid",
		})
		if err != nil {
			return
		}
		return
	}

	vars := r.URL.Query()
	id := vars.Get("id")
	url := "<https://api.coingecko.com/api/v3/coins?id=>" + id

	resp, err := http.Get(url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var coin *model.Token
	err = json.Unmarshal(body, &coin)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	h.coingeckoService.GetCoinInfo(coin)
}
