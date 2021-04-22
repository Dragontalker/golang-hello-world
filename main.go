package main

import (
	"log"
	"time"
)

type User struct {
	firstName   string
	lastName    string
	phoneNumber string
	age         int
	birthDate   time.Time
}

func main() {
	user := User{
		firstName:   "Richard",
		lastName:    "Yang",
		phoneNumber: "xxx-xxx-xxxx",
		age:         32,
	}

	log.Println(user.firstName, user.lastName, "Birthdate: ", user.birthDate)
}
