package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Task struct {
	gorm.Model
	Description string `gorm:"size:65536"`
	UserID      int
}

func (task Task) Speak(user User) string {
	return user.Name
}
