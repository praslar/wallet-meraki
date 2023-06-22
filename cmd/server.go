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
		DBName:     config.LoadEnv().DBName,
	})

	db = db.Debug()
	// error handling
	if err != nil {
		logrus.Errorf("Error connect db: %v", err.Error())
		return
	}

	repo := repo.NewRepo(db)

	authService := service.NewAuthService(repo)
	userService := service.NewUserService(repo, authService)
	userHandler := handler.NewUserHandler(userService, authService)

	tokenService := service.NewTokenService(repo)
	tokenHandler := handler.NewTokenHandler(tokenService, authService)

	migrateHandler := handler.NewMigrateHandler(db)
	r := mux.NewRouter()

	r.HandleFunc("/api/v1/register", userHandler.Register).Methods("POST")
	r.HandleFunc("/api/v1/login", userHandler.Login).Methods("POST")
	r.HandleFunc("/api/v1/user/get-all", userHandler.GetAllUser).Methods("GET")
	r.HandleFunc("/api/v1/crawl-data", tokenHandler.CrawlToken).Methods("POST")

	r.HandleFunc("/internal/migrate", migrateHandler.Migrate).Methods("POST")

	logrus.Infof("Start http server at :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		logrus.Errorf("Failed to start server, err: %v", err)
		return
	}
}
