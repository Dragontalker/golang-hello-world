package main

import "fmt"

func main() {
	whatToSay := saySomething("Hello, world!")
	fmt.Println(whatToSay)
}

func saySomething(s string) string {
	return s
}
