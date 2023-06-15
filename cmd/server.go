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
	if err != nil {
		logrus.Errorf("Error connecting to the database: %v", err.Error())
		return
	}

	userRepo := repo.NewUserRepo(db)

	authService := service.NewAuthService(userRepo)
	userService := service.NewUserService(userRepo, authService)

	authHandler := handler.NewAuthHandler(userService, authService)
	userHandler := handler.NewUserHandler(userService, authService)

	migrateHandler := handler.NewMigrateHandler(db)

	r := mux.NewRouter()

	apiRouter := r.PathPrefix("/api").Subrouter()

	// ADMIN routes
	adminRouter := apiRouter.PathPrefix("/v1/admin").Subrouter()
	adminRouter.Use(authHandler.AdminMiddleware)
	adminRouter.HandleFunc("/get-all", userHandler.GetAllUsers).Methods("GET")
	adminRouter.HandleFunc("/get-user/{userID}", userHandler.GetUser).Methods("GET")
	adminRouter.HandleFunc("/delete-user/{userID}", userHandler.DeleteUser).Methods("DELETE")
	adminRouter.HandleFunc("/update-role/{userID}", userHandler.UpdateUserRole).Methods("PUT")
	adminRouter.HandleFunc("/migrate", migrateHandler.Migrate).Methods("POST")

	r.HandleFunc("/register", userHandler.Register).Methods("POST")
	r.HandleFunc("/login", userHandler.Login).Methods("POST")

	logrus.Infof("Starting the server on port :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		logrus.Errorf("Failed to start the server: %v", err)
		return
	}
}
