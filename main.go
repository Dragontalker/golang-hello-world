package main

import (
	"log"	"time"
)

type User struct {
	firstName string
	lastName string
	phoneNUmber string
	age int
	birthDate time.Time
}

func main() {
	var s2 = "six"

	log.Println("s is", s)
	log.Println("s2 is", s2)

	saySomething("xxx")
}

func saySomething(s string) (string, string) {
	log.Println("s from saySomething is ", s)
	return s, "World"
}
