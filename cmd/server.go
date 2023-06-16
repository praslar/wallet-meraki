package main

import (
	"net/http"
	"wallet/internal/handler"
	"wallet/internal/middleware"
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

	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)
	migrateHandler := handler.NewMigrateHandler(db)
	r := mux.NewRouter()
	r.HandleFunc("/api/register", userHandler.Register).Methods("POST")
	r.HandleFunc("/api/login", userHandler.Login).Methods("POST")
	r.HandleFunc("/internal/migrate", migrateHandler.Migrate).Methods("POST")

	v1Group := r.PathPrefix("/api/v1").Subrouter()
	// Admin apis
	v1Group.HandleFunc("/admin/user/get-all", middleware.AuthenticateMiddleware(middleware.AuthorAdminMiddleware(userHandler.GetAllUsers))).Methods("GET")
	v1Group.HandleFunc("/admin/user/delete/{id}", middleware.AuthenticateMiddleware(middleware.AuthorAdminMiddleware(userHandler.DeleteUser))).Methods("DELETE")
	v1Group.HandleFunc("/admin/user/get-detail/{id}", middleware.AuthenticateMiddleware(middleware.AuthorAdminMiddleware(userHandler.GetUser))).Methods("GET")

	// User apis
	v1Group.HandleFunc("/user/get-info", middleware.AuthenticateMiddleware(userHandler.GetOne)).Methods("GET")
	if err := http.ListenAndServe(":8080", r); err != nil {
		logrus.Errorf("Failed to start server, err: %v", err)
		return
	}
}
