package handler

import (
	"net/http"
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
	// Đọc thông tin người dùng từ yêu cầu HTTP
	email := r.FormValue("email")
	password := r.FormValue("password")

	// Gọi hàm SignUp từ UserService
	err := h.userService.SignUp(email, password)
	if err != nil {
		// Xử lý lỗi
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Kiểm tra email đã đăng ký hay chưa
	exists, err := h.userService.CheckEmail(email)
	if err != nil {
		// Xử lý lỗi
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if exists {
		// Email đã được đăng ký, xử lý thông báo lỗi
		http.Error(w, "Email already registered", http.StatusBadRequest)
		return
	}

	// Xử lý thành công
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User registered successfully"))
}
