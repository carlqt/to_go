package main

import (
	"github.com/carlqt/to_go/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"strconv"
)

var db, err = gorm.Open("postgres", "dbname=to_go_dev sslmode=disable")

func errsToStr(err []error) (result []string) {
	for _, e := range err {
		result = append(result, e.Error())
	}
	return result
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:4567")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func main() {
	if err != nil {
		panic("Failed to connect to database")
	}

	router := gin.Default()
	router.Use(CORSMiddleware())

	router.GET("/", index)
	router.GET("/user/:id", show)
	router.GET("/tasks", taskIndex)
	router.GET("/user/:id/task", showUserTasks)
	router.POST("/test", test)
	router.POST("/user/:id/task", addUserTask)
	router.POST("/user", addUser)
	router.PUT("/task/:id", updateTask)
	router.DELETE("/destroy_all", destroyAll)

	router.Run(":9000")
}

func index(r *gin.Context) {
	var users []models.User
	db.Find(&users)
	r.JSON(200, users)
}

func addUser(r *gin.Context) {
	name := r.PostForm("name")
	age, err := strconv.Atoi(r.PostForm("age"))

	if err != nil {
		r.JSON(400, gin.H{"error": "Invalid age value"})

	} else {
		user := &models.User{Name: name, Age: age}

		if err := user.Validate(db); err != nil {
			r.JSON(400, gin.H{"errors": errsToStr(err)})
		} else {
			db.Create(&user)
			r.JSON(201, gin.H{"name": user.Name, "age": user.Age})
		}
	}
}

func show(r *gin.Context) {
	id, err := strconv.Atoi(r.Param("id"))

	if err != nil {
		r.JSON(400, gin.H{"error": "Invalid id value"})
		return
	}

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
	taskDescription := r.PostForm("description")

	user := db.Find(&models.User{}, id)

	if user.RecordNotFound() {
		r.JSON(404, gin.H{"error": "User not found"})
	} else {
		task := db.Create(&models.Task{Description: taskDescription, UserID: id})
		r.JSON(201, task.Value)
	}
}

func showUserTasks(r *gin.Context) {
	id, _ := strconv.Atoi(r.Param("id"))

	user := db.Find(&models.User{}, id)

	if user.RecordNotFound() {
		r.JSON(404, gin.H{"error": "User not found"})
	} else {
		var tasks []models.Task

		user.Related(&tasks)

		r.JSON(200, tasks)
	}
}

func updateTask(r *gin.Context) {
	var task models.Task
	id, _ := strconv.Atoi(r.Param("id"))
	desc := r.PostForm("description")

	if db.First(&task, id).RecordNotFound() {
		r.JSON(404, gin.H{"error": "User not found"})
	} else {
		db.Model(&task).Update("description", desc)
		r.JSON(200, task)
	}
}

func destroyAll(r *gin.Context) {
	db.Unscoped().Delete(&models.User{})
	db.Unscoped().Delete(&models.Task{})
}

func test(r *gin.Context) {
	desc := r.DefaultPostForm("description", "Failed")
	r.JSON(200, gin.H{"description": desc})
}
