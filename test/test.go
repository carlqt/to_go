package main

import (
	"bufio"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"os"
)

var logFile string = "logfile.log"

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

type User struct {
	Name string
	Age  int
}

func main() {
	fi, err := os.Open("input_file.txt")
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

	f, err := os.OpenFile(logFile, os.O_CREATE|os.O_APPEND, 0666)
	check(err, "fatal")
	writer := bufio.NewWriter(f)
	defer f.Close()

	fmt.Fprintln(writer, "Greed is good")
	writer.Flush()

	user := &User{"carl", "12"}
	fmt.Println(user)
}
