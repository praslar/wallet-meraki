package handler

import (
	"net/http"
	"wallet/internal/service"
)

type MigrateHandler struct {
	service service.MigrateService
}

func NewMigrateHandler(service service.MigrateService) MigrateHandler {
	return MigrateHandler{
		service: service}
}
func (h *MigrateHandler) Migrate(w http.ResponseWriter, r *http.Request) {
	h.service.Migrate(w, r)
}
