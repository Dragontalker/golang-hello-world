package main

import (
	"log"
)

func main() {
	mySlice := []string{"dog", "cat", "horse", "fish", "banana"}

	for _, x := range mySlice {
		log.Println(x)
	}
}
