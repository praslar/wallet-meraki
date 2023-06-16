package middleware

import (
	"encoding/json"
	"github.com/google/uuid"
	"net/http"
	"strings"
	"wallet/pkg/jwt"
)

func AuthenticateMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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
		claims, err := jwt.ValidJWTToken(token[1])
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"error": "unauthorized bearer",
			})
			return
		}
		// valid x-user-id
		xUserID, err := uuid.Parse(claims.XUserID)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"error": "unauthorized bearer",
			})
			return
		}
		r.Header.Set("x-user-id", xUserID.String())
		r.Header.Set("x-user-role", claims.XUserRole)
		next.ServeHTTP(w, r)
	}
}

func AuthorAdminMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		role := r.Header.Get("x-user-role")
		if role != "admin" {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"error": "don't have permission",
			})
			return
		}
		next.ServeHTTP(w, r)
		return
	}
}
