package handler

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"wallet/internal/model"
	"wallet/internal/service"
)

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) UserHandler {
	return UserHandler{
		userService: userService,
	}
}

func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	requestUser := model.UserRequest{}
	w.Header().Set("Content-Type", "application/json")

	err := json.NewDecoder(r.Body).Decode(&requestUser)
	if err != nil {
		logrus.Errorf("Failed to get request body : %v", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		err := json.NewEncoder(w).Encode(map[string]interface{}{
			"error": err.Error(),
		})
		if err != nil {
			logrus.Errorf(err.Error())
			return
		}
		return
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(requestUser.Password), bcrypt.DefaultCost)
	if err != nil {
		logrus.Errorf("Failed hashed password :%v", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		err := json.NewEncoder(w).Encode(map[string]interface{}{
			"error": err.Error(),
		})
		if err != nil {
			logrus.Errorf(err.Error())
			return
		}
	}
	if err := h.userService.Register(requestUser.Email, string(hashedPassword)); err != nil {
		logrus.Errorf("Failed to create user : %v", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		err := json.NewEncoder(w).Encode(map[string]interface{}{
			"error": err.Error(),
		})
		if err != nil {
			logrus.Errorf(err.Error())
			return
		}
	}
	if err = json.NewEncoder(w).Encode(requestUser); err != nil {
		logrus.Errorf("Failed to encode respone: %v", err.Error())
		return
	}
}
func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
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
			logrus.Errorf(err.Error())
			return
		}
		return
	}
	token, err := h.userService.Login(requestUser.Email, requestUser.Password)
	if err != nil {
		logrus.Errorf("Failed to get authenticate user : %v", err.Error())
		w.WriteHeader(http.StatusUnauthorized)
		err := json.NewEncoder(w).Encode(map[string]interface{}{
			"error": err.Error(),
		})
		if err != nil {
			logrus.Errorf(err.Error())
			return
		}
		return
	}
	response := map[string]interface{}{
		"token": token,
	}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		logrus.Errorf(err.Error())
		return
	}
}
