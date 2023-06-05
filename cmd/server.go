package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
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
		Host:     config.LoadEnv().Host,
		Port:     config.LoadEnv().Port,
		Username: config.LoadEnv().Username,
		Password: config.LoadEnv().Password,
		Dbname:   config.LoadEnv().Dbname,
	})
	// error handling
	if err != nil {
		fmt.Println("Đã có lỗi xảy ra: ", err)
		return
	}

	// khởi tạo user handler
	userRepo := repo.NewUserRepo(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	// Init the mux router
	router := mux.NewRouter()

	router.HandleFunc("/user/create", userHandler.Register).Methods("POST")

	// serve the app
	fmt.Println("Server at 5432")
	log.Fatal(http.ListenAndServe(":5432", router))

}
