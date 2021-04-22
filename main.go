package main

import "fmt"

func main() {
	fmt.Println(saySomething("Hello, world!"))
}

func saySomething(s string) string {
	return s
}
