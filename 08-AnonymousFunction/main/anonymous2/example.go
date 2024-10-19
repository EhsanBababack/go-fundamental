package main

import "fmt"

func main() {
	myFunc := func() {
		fmt.Println("Hello World")
	}

	myFunc()
	myFunc()
}