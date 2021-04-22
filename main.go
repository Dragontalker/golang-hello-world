package main

import (
	"log"
)

func main() {
	myMap := make(map[string]string)

	myMap["dog"] = "dog"
	myMap["fish"] = "fish"
	myMap["hat"] = "hat"

	for i, x := range myMap {
		log.Println(i, x)
	}
}
