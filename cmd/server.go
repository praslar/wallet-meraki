package main

import (
	"fmt"
	"wallet/config"
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
	fmt.Println(db)

	if err != nil {
		fmt.Println("Đã có lỗi xảy ra: ", err)
		return
	}
	fmt.Println(db)
	if err != nil {
		fmt.Println("Đã có lỗi xảy ra: ", err)
		return
	}
	fmt.Println(db)

	fmt.Println("check git")

}
