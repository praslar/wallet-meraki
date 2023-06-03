package main

import (
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
	"wallet/config"
	"wallet/internal/handler"
	"wallet/internal/repo"
	"wallet/internal/service"
	"wallet/pkg/pg"
)

func main() {

	config.SetEnv()
	db, err := pg.ConnectDB(config.AppConfig{
		DBHost:     config.LoadEnv().DBHost,
		DBPort:     config.LoadEnv().DBPort,
		DBUsername: config.LoadEnv().DBUsername,
		DBPassword: config.LoadEnv().DBPassword,
		DBDbname:   config.LoadEnv().DBDbname,
	})
	// error handling
	if err != nil {
		logrus.Errorf("Error connect db: %v", err.Error())
		return
	}
	r := mux.NewRouter()
	// MigrationDB
	migrateRepo := repo.NewMigrateRepo(db)
	migrateService := service.NewMigrateService(migrateRepo)
	migrateHandler := handler.NewMigrateHandler(migrateService)

	// UserLogin
	loginRepo := repo.NewUserRepo(db)
	loginService := service.NewUserService(loginRepo)
	loginHandler := handler.NewUserHandler(loginService)

	//  Define route
	r.HandleFunc("/internal/migrate", migrateHandler.Migrate).Methods("POST")
	r.HandleFunc("/user/login", loginHandler.UserLogin).Methods("GET")
	logrus.Infof("Start http server at :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		logrus.Errorf("Failed to start server, err: %v", err)
		return
	}

}
