package main

import (
	"fmt"
	"github.com/carlqt/to_go/models"
	"github.com/jinzhu/gorm"
)

func main() {
	db, err := gorm.Open("postgres", "dbname=to_go_dev sslmode=disable")
	if err != nil {
		panic("Failed to connect to database")
	}

	db.AutoMigrate(&models.User{})
	fmt.Println("Migration done")
	//var user models.User

	//db.Find(&user)
	//fmt.Println(user.Hello())
}
