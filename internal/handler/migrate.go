package handler

import (
	"gorm.io/gorm"
	"net/http"
	"wallet/internal/model"
)

type MigrateHandler struct {
	db *gorm.DB
}

func NewMigrateHandler(db *gorm.DB) MigrateHandler {
	return MigrateHandler{
		db: db,
	}
}

func (h *MigrateHandler) Migrate(w http.ResponseWriter, r *http.Request) {
	if err := h.db.AutoMigrate(&model.Role{}, &model.User{}); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
