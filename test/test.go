package main

import (
	"bufio"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"os"
)

func check(err error, level string) {
	if err != nil {
		switch level {
		case "error":
			log.Error(err)
		case "info":
			log.Info(err)
		case "fatal":
			log.Fatal(err)
		default:
			panic("Unrecognized error level")
		}
	}
}

func main() {
	fi, err := os.Open("input_fil.txt")
	check(err, "fatal")

	defer func() {
		log.Info("closing file")
		fi.Close()
	}()

	scanner := bufio.NewScanner(fi)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	check(scanner.Err(), "error")
}
