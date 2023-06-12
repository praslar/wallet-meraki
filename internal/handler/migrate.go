package handler

import (
	"encoding/json"
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
	if err := h.db.AutoMigrate(&model.User{}, &model.Wallet{}, &model.Role{}, &model.Token{}, &model.Transaction{}); err != nil {
		json.NewEncoder(w).Encode(map[string]interface{}{
			"error": err.Error(),
		})
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
