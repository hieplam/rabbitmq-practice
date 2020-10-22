package utils

import (
	"log"
	"os"
	"strings"
)

func FailOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func BodyFrom(args []string) string {
	var s string
	if len(args) < 2 || os.Args[1] == "" {
		s = "Hello"
	} else {
		s = strings.Join(args[1:], " ")
	}
	return s
}
