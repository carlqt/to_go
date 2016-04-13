package main

import (
	"fmt"
	"github.com/carlqt/to_go/models"
)

type Rectangle struct {
	Length int
	Width  int
}

func main() {
	u := models.User{"Carl", 15}
	fmt.Println(u)
}
