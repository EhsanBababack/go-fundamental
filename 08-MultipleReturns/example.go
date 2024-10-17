package main

import "fmt"

func main(){
	fmt.Println(CalculateRoomPrice("single", 3, 2))
    fmt.Println(CalculateRoomPrice("double", 4, 3))
    fmt.Println(CalculateRoomPrice("suite", 2, 4))
	order1,tax1  := CalculateRoomPrice("single", 5, 2)
	fmt.Printf("order 4 is : %d , tax 4 is : %f" , order1 , tax1)
}

func CalculateRoomPrice(roomType string , nights int , personCount int)(int , float64){
	var price int
	var tax float64
	switch roomType{
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
	return price , tax
}