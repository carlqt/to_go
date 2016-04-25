package models

import (
	"errors"
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

func (u *User) Validate(db *gorm.DB) (err []error) {
	if u.Age <= 0 {
		err = append(err, errors.New("Age should be greater than 0"))
	}

	if !db.Where("name = ?", u.Name).Find(&User{}).RecordNotFound() {
		err = append(err, errors.New("User already exists"))
	}

	return err
}
