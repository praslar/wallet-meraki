package handler

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"wallet/internal/model"
	"wallet/internal/service"
	"wallet/internal/utils"

	"github.com/sirupsen/logrus"
)

type UserHandler struct {
	userService service.UserService
	authService service.AuthService
}

func NewUserHandler(userService service.UserService, authService service.AuthService) *UserHandler {
	return &UserHandler{
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
		json.NewEncoder(w).Encode(map[string]interface{}{
			"error": err.Error(),
		})
		return
	}

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
		logrus.Errorf("Failed to create user: %v", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"error": err.Error(),
		})
		return
	}

	json.NewEncoder(w).Encode(requestUser)
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
		logrus.Errorf("Failed to hash password: %v", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"error": err.Error(),
		})
		return
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

	json.NewEncoder(w).Encode(map[string]interface{}{
		"token": token,
	})
}

func (h *UserHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	// Parse pagination parameters
	page, err := strconv.Atoi(r.FormValue("page"))
	if err != nil {
		page = 1
	}

	limit, err := strconv.Atoi(r.FormValue("limit"))
	if err != nil || limit <= 0 {
		limit = 10
	}

	// Parse sorting parameters
	sortField := r.FormValue("sortField")
	sortOrder := r.FormValue("sortOrder")

	// Parse filtering parameter
	filterName := r.FormValue("filterName")

	users, totalPages, err := h.userService.GetAllUsers(page, limit, sortField, sortOrder, filterName)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"error": err.Error(),
		})
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"data": users,
		"meta": map[string]interface{}{
			"totalPages": totalPages,
		},
	})
}
func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userIDStr := params["userID"]
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"error": "invalid user ID",
		})
		return
	}

	user, err := h.userService.GetUser(userID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"error": "failed to get user",
		})
		return
	}

	json.NewEncoder(w).Encode(user)
}

func (h *UserHandler) UpdateUserRole(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userIDStr := params["userID"]
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"error": "invalid user ID",
		})
		return
	}

	err = h.userService.UpdateUserRole(userID, "admin") // Set the new role as "admin"
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"error": "failed to update user role",
		})
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "user role updated successfully",
	})
}
