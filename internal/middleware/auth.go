package middleware

import (
	"encoding/json"
	"net/http"
	"strings"
)

func AuthenMiddleware(next http.Handler) http.Handler {
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
		next.ServeHTTP(w, r)
	})
}