package main

import "fmt"

func main() {
	price1, tax1, finalPrice1 := CalculateRoomPrice("single", 2, 3)
	fmt.Printf("order 1 - price is : %d - tax is %f ---> final price is : %d" , price1 , tax1 , finalPrice1)
}

func CalculateRoomPrice(roomType string, nights int, personCount int) (price int, tax float64, finalPrice int) {
	switch roomType {
	case "single":
		price = 100 * nights * personCount
		tax = 0.08 * float64(price)
	case "double":
		price = 150 * nights * personCount
		tax = 0.12 * float64(price)
	case "suite":
		price = 200 * nights * personCount
		tax = 0.15 * float64(price)
	}
	finalPrice = price + int(tax)
	return
}