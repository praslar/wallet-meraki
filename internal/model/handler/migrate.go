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
<<<<<<< HEAD
	if err := h.db.AutoMigrate(&model.User{}, &model.Role{}, &model.Wallet{}, &model.Token{}); err != nil {
		err := json.NewEncoder(w).Encode(map[string]interface{}{
=======
	if err := h.db.AutoMigrate(&model.User{}, &model.Wallet{}, &model.Role{}, &model.Token{}, &model.Transaction{}); err != nil {
		json.NewEncoder(w).Encode(map[string]interface{}{
>>>>>>> f42f72261765b586a57e931f5a776a40c861c8d0
			"error": err.Error(),
		})
		if err != nil {
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
