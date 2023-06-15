package handler

import (
	"encoding/json"
	"net/http"
	"strings"
	"wallet/internal/service"

	"github.com/sirupsen/logrus"
)

type AuthHandler struct {
	userService service.UserService
	authService service.AuthService
}

func NewAuthHandler(userService service.UserService, authService service.AuthService) AuthHandler {
	return AuthHandler{
		userService: userService,
		authService: authService,
	}
}

func (ah *AuthHandler) AdminMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")

		jwtToken := r.Header.Get("Authorization")
		token := strings.Split(jwtToken, " ")
		if len(token) != 2 || token[0] != "Bearer" {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"error": "unauthorized bearer",
			})
			return
		}

		err := ah.authService.ValidJWTToken(token[1], "admin")
		if err != nil {
			logrus.Errorf("Failed to validate token: %v", err.Error())
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"error": "unauthorized token",
			})
			return
		}

		// Call the next handler in the chain
		next.ServeHTTP(w, r)
	})
}
