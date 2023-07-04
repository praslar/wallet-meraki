package handler

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"wallet/internal/model"
	"wallet/internal/service"
	"wallet/internal/utils"
)

type UserHandler struct {
	userService   service.UserService
	walletService service.WalletServiceInterface
}

func NewUserHandler(userService service.UserService, walletService service.WalletServiceInterface) UserHandler {
	return UserHandler{
		userService:   userService,
		walletService: walletService,
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

func (h *UserHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	// Parse pagination parameters
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		page = 1
	}

	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil || limit <= 0 {
		limit = 10
	}

	// Parse filtering parameter
	filterName := r.URL.Query().Get("email")

	// Parse sorting parameter
	sortOrder := r.URL.Query().Get("sort")

	users, totalPages, err := h.userService.GetAllUsers(filterName, sortOrder, page, limit)
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
		return
	}
}

func (h *UserHandler) ViewTransaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	currentUser := r.Header.Get("x-user-id")
	users, err := h.userService.GetTransactionID(currentUser)
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
		return
	}
	return
}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userIDStr := params["id"]
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

func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userIDStr := params["id"]
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"error": "invalid user ID",
		})
		return
	}

	err = h.userService.DeleteUser(userID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"error": "failed to delete user",
		})
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "user deleted successfully",
	})
}

func (h *UserHandler) UpdateUserRole(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	userIDStr := params["id"]
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

	return
}

func (h *UserHandler) GetListAllTransaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
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
