package repo

import (
	"gorm.io/gorm"
	"net/http"
	"wallet/internal/model"
)

type MigrateRepo struct {
	db *gorm.DB
}

func NewMigrateRepo(db *gorm.DB) MigrateRepo {
	return MigrateRepo{
		db: db,
	}
}
func (rp *MigrateRepo) Migrate(w http.ResponseWriter, r *http.Request) {
	if err := rp.db.AutoMigrate(&model.User{}, &model.Role{}); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
