package main

import "fmt"

func main() {
	var whatToSay string
	whatToSay = saySomething("Hello, world!")
	fmt.Println(whatToSay)
}

func saySomething(s string) string {
	return s
}
