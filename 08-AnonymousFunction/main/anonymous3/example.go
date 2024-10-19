package main

import "fmt"

func main() {
	fmt.Println(func (numbers ...int)int  {
		var total int
		for _,number := range numbers{
			total+=number
		}
		return total
	}(1,2,6,5,8,9,7,6))
}