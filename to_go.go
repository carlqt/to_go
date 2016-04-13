package main

import (
	"./models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func main() {
	router := gin.Default()

	db, err := gorm.Open("postgres", "dbname=to_go_dev sslmode=disable")
	if err != nil {
		panic("Failed to connect to database")
	}

	router.GET("/", index)

	router.Run(":9000")
}

func index(r *gin.Context) {
	var user models.User
	db.First(&user)
	r.JSON(200, gin.H{"name": ""})
}
