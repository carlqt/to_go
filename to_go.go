package main

import (
	"fmt"
	"github.com/carlqt/to_go/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"strconv"
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
	var users []models.User
	db.Find(&users)
	r.JSON(200, users)
}

func addUser(r *gin.Context) {

	name := r.DefaultQuery("name", "Anonymous")
	age, _ := strconv.Atoi(r.Query("age"))

	db.Create(&models.User{Name: name, Age: age})

	r.JSON(201, gin.H{"name": name, "age": age})
}

func show(r *gin.Context) {
	id := r.Param("id")

	user := db.Find(&models.User{}, id)

	fmt.Println(user)
	r.JSON(200, user.Value)
}
