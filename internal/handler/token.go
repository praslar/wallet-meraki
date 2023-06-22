package handler

import (
	"encoding/json"
	"net/http"
	"wallet/internal/service"
)

type TokenHandler struct {
	tokenService service.TokenService
	authService  service.AuthService
}

func NewTokenHandler(tokenService service.TokenService, authService service.AuthService) TokenHandler {
	return TokenHandler{
		tokenService: tokenService,
		authService:  authService,
	}
}

func (h *TokenHandler) CrawlToken(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	err := h.tokenService.TriggerCrawl()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"error": err.Error(),
		})
		return
	}
}
