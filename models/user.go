package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type User struct {
	gorm.Model
	Name string
	Age  int
}

func (user User) Hello() string {
	content := fmt.Sprintf("Hello, my name is %s and I am %d years old", user.Name, user.Age)
	return content
}
