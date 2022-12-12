package main

import (
	"fmt"

	"gx_myfood/Config"
	"gx_myfood/Routes"

	"github.com/jinzhu/gorm"
)

var err error

func main() {
	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
	defer Config.DB.Close()

	if err != nil {
		fmt.Println("Status Error: ", err)
	}

	r := Routes.SetupRoutes()
	r.Run(":3000")
}
