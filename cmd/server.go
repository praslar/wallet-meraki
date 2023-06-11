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
	walletService := service.NewWalletService(userRepo)
	tokenService := service.NewTokenService(userRepo)

	// userHandler
	userHandler := handler.NewUserHandler(userService, authService, walletService, tokenService)
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
	//r.HandleFunc("/api/v1/admin/wallet/create/token", userHandler.CreateTokenAd).Methods("POST")
	//r.HandleFunc("/api/v1/admin/wallet/update/token", userHandler.UpdateTokenAd).Methods("PUT")
	//r.HandleFunc("/api/v1/admin/wallet/delete/token", userHandler.DeleteTokenAd).Methods("DELETE")
	//r.HandleFunc("/api/v1/admin/wallet/transfer/token", userHandler.TransferTokenAd).Methods("POST")

	logrus.Infof("Start http server at :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		logrus.Errorf("Failed to start server, err: %v", err)
		return
	}
}
