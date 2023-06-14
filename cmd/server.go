package main

import (
	"net/http"
	"wallet/internal/handler"
	"wallet/internal/repo"
	"wallet/internal/service"
	"wallet/pkg"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"wallet/config"
)

func main() {

	config.SetEnv()

	db, err := pkg.ConnectDB(config.AppConfig{
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
	walletService := service.NewWalletService(userRepo)
	tokenService := service.NewTokenService(userRepo)
	coingeckoService := service.NewCoingeckoService(userRepo)

	// userHandler
	userHandler := handler.NewUserHandler(userService, authService, walletService, tokenService, coingeckoService)
	// migrateHandler
	migrateHandler := handler.NewMigrateHandler(db)

	r := mux.NewRouter()

	r.HandleFunc("/internal/migrate", migrateHandler.Migrate).Methods("POST")
	r.HandleFunc("/api/v1/register", userHandler.Register).Methods("POST")
	r.HandleFunc("/api/v1/login", userHandler.Login).Methods("POST")
	r.HandleFunc("/api/v1/admin/get-all-user", userHandler.GetAllUser).Methods("GET")
	//Wallet
	r.HandleFunc("/api/v1/user/wallet/create", userHandler.CreateWallet).Methods("POST")
	//Admin-TokenServices
	r.HandleFunc("/api/v1/admin/create/token", userHandler.CreateToken).Methods("POST")
	r.HandleFunc("/api/v1/admin/delete/token", userHandler.DeleteToken).Methods("DELETE")
	r.HandleFunc("/api/v1/admin/update/token", userHandler.UpdateToken).Methods("PUT")
	r.HandleFunc("/api/v1/admin/transfer/token", userHandler.SendUserToken).Methods("POST")
	//Crawl Data List All Coin
	r.HandleFunc("/coins/{id}", userHandler.GetCoinInfo).Methods("GET")
	http.Handle("/", r)

	logrus.Infof("Start http server at :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		logrus.Errorf("Failed to start server, err: %v", err)
		return
	}
}
