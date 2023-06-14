package handler

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"wallet/internal/model"
	"wallet/internal/service"
)

type UserHandler struct {
	userService service.UserService
	authService service.AuthService
}

func NewUserHandler(userService service.UserService, authService service.AuthService) UserHandler {
	return UserHandler{
		userService: userService,
		authService: authService,
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

	_, err := h.authService.ValidJWTToken(token[1], "user")
	if err != nil {
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

func (h *UserHandler) GetTransactionID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params, _ := url.ParseQuery(r.URL.RawQuery)
	id := params["id"][0]
	id = url.QueryEscape(id)
	result, err := h.userService.GetTransactionID(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{"error": err.Error()})
		return
	}
	if err = json.NewEncoder(w).Encode(result); err != nil {
		return
	}

}

func (h *UserHandler) ViewTransaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	headers := r.Header
	authHeader := headers.Get("Authorization")
	tokenString := strings.Split(authHeader, " ")

	claims, err := h.authService.ValidJWTToken(tokenString[1], "user")
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		err := json.NewEncoder(w).Encode(map[string]interface{}{
			"error": "fail authorized",
		})
		if err != nil {
			return
		}
		return
	}

	result, err := h.userService.GetTransactionID(claims.XUserID)
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
		"data": result,
	}); err != nil {
		return
	}

}

func (h *UserHandler) GetListAllTransaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	headers := r.Header
	authHeader := headers.Get("Authorization")
	tokenString := strings.Split(authHeader, " ")

	_, err := h.authService.ValidJWTToken(tokenString[1], "admin")
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		err := json.NewEncoder(w).Encode(map[string]interface{}{
			"error": "fail authorized",
		})
		if err != nil {
			return
		}
		return
	}
	//su dung ham take trong mux
	fromWallet := r.URL.Query().Get("from_wallet")
	toWallet := r.URL.Query().Get("to_wallet")
	tokenAddress := r.URL.Query().Get("token_address")
	amount := r.URL.Query().Get("amount")
	amountInt, _ := strconv.Atoi(amount)
	email := r.URL.Query().Get("email")
	orderBy := r.URL.Query().Get("order_by")
	//pagesize,page
	pageSize := r.URL.Query().Get("page_size")
	pageSizeInt, _ := strconv.Atoi(pageSize)
	page := r.URL.Query().Get("page")
	pageInt, _ := strconv.Atoi(page)

	result, err := h.userService.GetTransaction(fromWallet, toWallet, email, tokenAddress, orderBy, amountInt, pageSizeInt, pageInt)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{"error": err.Error()})
		return
	}
	if err = json.NewEncoder(w).Encode(map[string]interface{}{
		"data": result,
	}); err != nil {
		return
	}
}
