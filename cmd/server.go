package main

import (
	"net/http"
	"wallet/internal/handler"
	"wallet/internal/model"
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

	// error handling
	if err != nil {
		logrus.Errorf("Error connect db: %v", err.Error())
		return
	}
	logrus.Infof("Connect db successfull. Database name: %s", db.Name())
	logrus.Infof("Start http server at :8080")

	userRepo := repo.NewUserRepo(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)
	r := mux.NewRouter()

	r.HandleFunc("/register", userHandler.Register).Methods("POST")
	r.HandleFunc("/migrate", Migrate).Methods("GET")

	logrus.Infof("API register: %s", "/register POST")
	logrus.Infof("API migrate: %s", "/migrate GET")

	if err := http.ListenAndServe(":8080", r); err != nil {
		logrus.Errorf("Failed to start server, err: %v", err)
		return
	}
}

func Migrate(w http.ResponseWriter, r *http.Request) {
	config.SetEnv()

	db, err := pg.ConnectDB(config.AppConfig{
		DBHost:     config.LoadEnv().DBHost,
		DBPort:     config.LoadEnv().DBPort,
		DBUsername: config.LoadEnv().DBUsername,
		DBPassword: config.LoadEnv().DBPassword,
		Dbname:     config.LoadEnv().Dbname,
	})
	if err != nil {
		return
	}
	db.AutoMigrate(&model.User{})
}
