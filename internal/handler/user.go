package handler

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"net/http"
	"wallet/internal/model"
	"wallet/internal/service"
	"wallet/internal/utils"
)

type UserHandler struct {
	userService  service.UserService
	tokenService service.TokenService
}

func NewUserHandler(userService service.UserService, tokenService service.TokenService) UserHandler {
	return UserHandler{
		userService:  userService,
		tokenService: tokenService,
	}
}

func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	requestUser := model.UserRequest{}
	w.Header().Set("Content-Type", "application/json")

	err := json.NewDecoder(r.Body).Decode(&requestUser)
	if err != nil {
		logrus.Errorf("Failed to get request body: %v", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"error": err.Error(),
		})
		return
	}
	//hashpw
	hashedPassword, err := utils.HashPassword(requestUser.Password)
	if err != nil {
		logrus.Errorf("Failed to hash password: %v", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"error": err.Error(),
		})
		return
	}

	if err := h.userService.Register(requestUser.Email, hashedPassword); err != nil {
		logrus.Errorf("Failed create user: %v", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"error": err.Error(),
		})
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
		json.NewEncoder(w).Encode(map[string]interface{}{
			"error": err.Error(),
		})
		return
	}

	hashedPassword, err := utils.HashPassword(requestUser.Password)
	if err != nil {
		logrus.Errorf("Fail to hash password: %v", err.Error())
	}

	token, err := h.userService.Login(requestUser.Email, hashedPassword)
	if err != nil {
		logrus.Errorf("Failed to login: %v", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"error": "failed to login",
		})
		return
	}

	if err = json.NewEncoder(w).Encode(map[string]interface{}{
		"token": token,
	}); err != nil {
		return
	}
}

func (h *UserHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	//jwtToken := r.Header.Get("Authorization")
	//token := strings.Split(jwtToken, " ")
	//if token[0] != "Bearer" {
	//	w.WriteHeader(http.StatusUnauthorized)
	//	json.NewEncoder(w).Encode(map[string]interface{}{
	//		"error": "unauthorized",
	//	})
	//	return
	//}
	//
	//// jwtToken
	//if err := h.authService.ValidJWTToken(token[1], "admin"); err != nil {
	//	w.WriteHeader(http.StatusUnauthorized)
	//	json.NewEncoder(w).Encode(map[string]interface{}{
	//		"error": "unauthorized",
	//	})
	//	return
	//}

	name := r.URL.Query().Get("name")
	fmt.Println(name)

	orderBy := r.URL.Query().Get("sort")
	users, err := h.userService.GetAllUser(orderBy)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"error": err.Error(),
		})
		return
	}
	if err = json.NewEncoder(w).Encode(map[string]interface{}{
		"data": users,
	}); err != nil {
		return
	}
}

func (h *UserHandler) CreateToken(w http.ResponseWriter, r *http.Request) {
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
	}
}

func (h *UserHandler) GetOne(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	currentUser := r.Header.Get("x-user-id")
	users, err := h.userService.GetUserByID(currentUser)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logrus.Errorf(err.Error())
		json.NewEncoder(w).Encode(map[string]interface{}{
			"error": "Internal server error",
		})
		return
	}

	if err := json.NewEncoder(w).Encode(map[string]interface{}{
		"data": users,
	}); err != nil {
		//>>>>>>> develop
		return
	}
}
