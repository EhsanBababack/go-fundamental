package main

import "fmt"

func main() {
	sum1, mul1 := Calculator(1, 2, 3, 6, 5, 4, 8, 9, 6, 23, 2, 5, 6, 4)
	fmt.Printf("sum is : %d , multiply is : %d" , sum1 , mul1)
}

func Calculator(numbers ...int) (sum int, mul int) {
	mul = 1
	for _, number := range numbers {
		sum += number
		mul *= number
	}
	return
}