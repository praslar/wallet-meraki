package handler

import (
	"encoding/json"
	"net/http"
	"wallet/internal/model"
	"wallet/internal/service"
)

type UserHandler struct {
	service service.UserService
}

func NewUserHandler(service service.UserService) UserHandler {
	return UserHandler{service: service}
}

// ---------------------- responseWithJSON  ----------------------

func responseWithJSON(response http.ResponseWriter, statusCode int, data interface{}) {
	result, _ := json.Marshal(data)
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(statusCode)
	_, err := response.Write(result)
	if err != nil {
		return
	}
}

// ---------------------- responseWithError  ----------------------

func responseWithError(response http.ResponseWriter, statusCode int, msg string) {
	responseWithJSON(response, statusCode, map[string]string{
		"error": msg,
	})
}

func (h *UserHandler) CreateToken() string {
	return h.service.CreateToken()
}

func (h *UserHandler) UserLogin(w http.ResponseWriter, r *http.Request) {

	var user *model.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		responseWithError(w, http.StatusBadRequest, err.Error())
	}
	email := user.Email
	password := user.Password
	user, err = h.service.UserLogin(email, password)
	if err != nil {
		responseWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	token := h.CreateToken()
	responseWithJSON(w, http.StatusOK, token)

}
