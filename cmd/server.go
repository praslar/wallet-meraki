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
	// userRepo
	userRepo := repo.NewUserRepo(db)

	// userService
	authService := service.NewAuthService(userRepo)
	userService := service.NewUserService(userRepo)
	tokenService := service.NewTokenService(userRepo)

	// userHandler
	userHandler := handler.NewUserHandler(userService, authService, tokenService)
	// migrateHandler
	migrateHandler := handler.NewMigrateHandler(db)

	r := mux.NewRouter()

	r.HandleFunc("/internal/migrate", migrateHandler.Migrate).Methods("POST")
	r.HandleFunc("/api/v1/register", userHandler.Register).Methods("POST")
	r.HandleFunc("/api/v1/login", userHandler.Login).Methods("POST")
	r.HandleFunc("/api/v1/admin/get-all-user", userHandler.GetAllUser).Methods("GET")

	//Admin-TokenServices
	r.HandleFunc("/api/v1/admin/create/token", userHandler.CreateToken).Methods("POST")
	r.HandleFunc("/api/v1/admin/delete/token", userHandler.DeleteToken).Methods("DELETE")
	r.HandleFunc("/api/v1/admin/update/token", userHandler.UpdateToken).Methods("PUT")
	//Connect to http server
	logrus.Infof("Start http server at :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		logrus.Errorf("Failed to start server, err: %v", err)
		return
	}
}
