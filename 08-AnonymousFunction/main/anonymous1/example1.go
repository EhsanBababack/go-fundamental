package main

import "fmt"

func main() {
	func() {
		fmt.Printf("Hello World")
	}()
}