package main

import "fmt"

func main() {
	myFunc := func (numbers ...int)int  {
		var total int
		for _,number := range numbers{
			total+=number
		}
		return total
	}
	 
	fmt.Println(myFunc(1, 2, 3, 6, 5, 4, 8))
}