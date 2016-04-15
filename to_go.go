package main

import (
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
	router.GET("/user/:id", show)
	router.GET("/tasks", taskIndex)
	router.GET("/test", test)
	router.POST("/user/:id/task", addUserTask)
	router.POST("/user", addUser)

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
	if user.RecordNotFound() {
		r.JSON(404, gin.H{"error": "Record not found"})
	} else {
		r.JSON(200, user.Value)
	}
}

func taskIndex(r *gin.Context) {
	var tasks []models.Task
	db.Find(&tasks)

	r.JSON(200, tasks)
}

func addUserTask(r *gin.Context) {
	id, _ := strconv.Atoi(r.Param("id"))
	taskDescription := r.Query("description")

	user := db.Find(&models.User{}, id)

	if user.RecordNotFound() {
		r.JSON(404, gin.H{"error": "User not found"})
	} else {
		task := db.Create(&models.Task{Description: taskDescription, UserID: id})
		r.JSON(201, task)
	}
}

func test(r *gin.Context) {
	desc := r.DefaultQuery("description", "Failed")
	r.JSON(200, gin.H{"description": desc})
}
