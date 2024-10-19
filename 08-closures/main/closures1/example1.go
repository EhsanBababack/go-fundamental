package main

func main() {
	firstName := "Ali"
	printFirstNameFunc := func ()  {
		println(firstName)
	}

	firstName = "Reza"
	printFirstNameFunc()
}