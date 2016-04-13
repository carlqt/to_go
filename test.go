package main

import (
	"./models"
	"fmt"
)

type Rectangle struct {
	Length int
	Width  int
}

func main() {
	u := models.User{"Carl", 15}
	fmt.Println(u)
}
