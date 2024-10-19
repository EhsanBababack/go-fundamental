package main

func main() {
	defer println("Bye")
	defer println("Good")
	defer println("Morning")
	for i := 0; i < 10; i++ {
		defer println(i)
	}
	println("Hello")
}