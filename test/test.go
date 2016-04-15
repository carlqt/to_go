package main

import (
	"fmt"
	"github.com/carlqt/to_go/models"
	"github.com/jinzhu/gorm"
)

var db, err = gorm.Open("postgres", "dbname=to_go_dev sslmode=disable")

func main() {
	var tasks []models.Task
	var user models.User

	db.Find(&user, 1).Related(&tasks)

	fmt.Println(tasks)
}
