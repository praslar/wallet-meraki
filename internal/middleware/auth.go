package middleware

import (
	"encoding/json"
	"fmt"
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

func AuthourMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		jwtToken := r.Header.Get("x-user-id")
		// TODO:
		// get user current roles
		// check if user has role admin
		// if not, return unauthorized
		fmt.Println(jwtToken)
		next.ServeHTTP(w, r)
	})
}
