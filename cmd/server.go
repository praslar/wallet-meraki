package main

import (
	"net/http"
	"wallet/config"
	"wallet/internal/handler"
	"wallet/internal/repo"
	"wallet/internal/service"
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

	authService := service.NewAuthService(userRepo)
	userService := service.NewUserService(userRepo, authService)
	userHandler := handler.NewUserHandler(userService, authService)
	migrateHandler := handler.NewMigrateHandler(db)

	walletRepo := repo.NewWalletRepo(db)
	walletService := service.NewWalletService(walletRepo, authService)
	walletHandler := handler.NewWalletHandler(walletService, authService)

	r := mux.NewRouter()
	//User
	r.HandleFunc("/api/v1/register", userHandler.Register).Methods("POST")
	r.HandleFunc("/api/v1/login", userHandler.Login).Methods("POST")
	r.HandleFunc("/api/v1/user/get-all", userHandler.GetAllUser).Methods("GET")
	//Wallet
	r.HandleFunc("/api/v1/wallet/create", walletHandler.CreateWallet).Methods("POST")
	r.HandleFunc("/api/v1/wallet/get-all", walletHandler.GetAllWallet).Methods("GET")
	r.HandleFunc("/api/v1/wallet/update-wallet", walletHandler.UpdateWallet).Methods("PUT")
	r.HandleFunc("/api/v1/wallet/delete-wallet", walletHandler.DeleteWallet).Methods("DELETE")
	//Migrate
	r.HandleFunc("/internal/migrate", migrateHandler.Migrate).Methods("POST")

	logrus.Infof("Start http server at :8000")
	if err := http.ListenAndServe(":8000", r); err != nil {
		logrus.Errorf("Failed to start server, err: %v", err)
		return
	}
}
