package main

import "log"

func main() {
	var mySlice []string

	mySlice = append(mySlice, "Richard")
	mySlice = append(mySlice, "John")

	log.Println(mySlice)
}
