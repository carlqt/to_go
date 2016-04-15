package main

import (
	"fmt"
)

type Rectangle struct {
	Length int
	Width  int
}

func init() {
	u := Rectangle{14, 15}
	fmt.Println(u)
}
