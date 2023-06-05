package main

import (
	"net/http"
	"wallet/internal/handler"
	"wallet/internal/repo"
	"wallet/internal/service"

	"wallet/config"
	"wallet/pkg/pg"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

func main() {

	config.SetEnv()

	db, err := pg.ConnectDB(config.AppConfig{
		DBHost:     config.LoadEnv().DBHost,
		DBPort:     config.LoadEnv().DBPort,
		DBUsername: config.LoadEnv().DBUsername,
		DBPassword: config.LoadEnv().DBPassword,
		Dbname:     config.LoadEnv().Dbname,
	})

	db = db.Debug()
	// error handling
	if err != nil {
		logrus.Errorf("Error connect db: %v", err.Error())
		return
	}

	userRepo := repo.NewUserRepo(db)

	authService := service.NewAuthService()
	userService := service.NewUserService(userRepo, authService)
	userHandler := handler.NewUserHandler(userService)
	migrateHandler := handler.NewMigrateHandler(db)
	r := mux.NewRouter()

	r.HandleFunc("/api/v1/register", userHandler.Register).Methods("POST")
	r.HandleFunc("/api/v1/login", userHandler.Login).Methods("POST")

	r.HandleFunc("/internal/migrate", migrateHandler.Migrate).Methods("POST")

	logrus.Infof("Start http server at :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		logrus.Errorf("Failed to start server, err: %v", err)
		return
	}
}
