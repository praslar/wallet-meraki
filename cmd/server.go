package main

import (
	"fmt"
	"wallet/config"
	"wallet/pkg/pg"
)

func main() {
	db, err := pg.ConnectDB(config.AppConfig{
		Host:     "172.104.41.46",
		Port:     "5432",
		Username: "nmadmin",
		Password: "L9nLshJ3F7NqJgu7",
		Dbname:   "meraki",
	})
	// error handling
	if err != nil {
		fmt.Println("Đã có lỗi xảy ra: ", err)
		return
	}
	fmt.Println(db)
}
