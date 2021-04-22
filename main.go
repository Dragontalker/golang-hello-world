package main

import "log"

func main() {
	myMap := make(map[string]string)

	myMap["dog"] = "Samson"

	log.Println(myMap["dog"])
}
