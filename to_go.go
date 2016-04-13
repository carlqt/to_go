package main

import (
	"fmt"
	"github.com/carlqt/to_go/models"
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "reflect"
)

var db, err = gorm.Open("postgres", "dbname=to_go_dev sslmode=disable")

func main() {
	if err != nil {
		panic("Failed to connect to database")
	}

	router := gin.Default()

	router.GET("/", index)
	router.POST("/user", addUser)
	router.GET("/user/:id", show)

	router.Run(":9000")
}

func index(r *gin.Context) {
	var user models.User
	db.Last(&user)
	r.JSON(200, gin.H{"name": user.Name})
}

func addUser(r *gin.Context) {
	user := models.User{Name: "Anne", Age: 24}

	db.Create(&user)
	r.JSON(201, gin.H{"success": "Yeah"})
}

func show(r *gin.Context) {
	id := r.Param("id")

	user := db.Find(&models.User{}, id)

	js := structs.Map(user)

	fmt.Println(user)
	r.JSON(200, js["Value"])
}
