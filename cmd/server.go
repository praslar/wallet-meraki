package main

import (
	"fmt"
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
	// userRepo
	userRepo := repo.NewUserRepo(db)
<<<<<<< HEAD

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

=======

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
>>>>>>> f42f72261765b586a57e931f5a776a40c861c8d0
	r.HandleFunc("/internal/migrate", migrateHandler.Migrate).Methods("POST")
	r.HandleFunc("/api/v1/register", userHandler.Register).Methods("POST")
	r.HandleFunc("/api/v1/login", userHandler.Login).Methods("POST")
	r.HandleFunc("/api/v1/admin/get-all-user", userHandler.GetAllUser).Methods("GET")
	//Wallet
	r.HandleFunc("/api/v1/user/wallet/create", userHandler.CreateWallet).Methods("POST")
	//Admin-TokenServices
	r.HandleFunc("/api/v1/admin/create/token", userHandler.CreateToken).Methods("POST")
	r.HandleFunc("/api/v1/admin/update/token", userHandler.UpdateToken).Methods("PUT")
	r.HandleFunc("/api/v1/admin/delete/token", userHandler.DeleteToken).Methods("DELETE")
	//r.HandleFunc("/api/v1/admin/wallet/transfer/token", userHandler.TransferTokenAd).Methods("POST")

	logrus.Infof("Start http server at :8000")
	if err := http.ListenAndServe(":8000", r); err != nil {
		logrus.Errorf("Failed to start server, err: %v", err)
		return
	}
	fmt.Print("test2")
}
