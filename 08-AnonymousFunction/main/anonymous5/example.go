package main

import (
	"fmt"
	"sort"
)

func main() {
	numbers := []int{12,5,65,89,2,3,45,17,69,32,69,87,56,21,11}
	fmt.Printf("%v \n" , numbers)

	sort.Slice(numbers , func(i, j int) bool {
		return numbers[i]<numbers[j]
	})
	fmt.Printf("%v \n" , numbers)
}