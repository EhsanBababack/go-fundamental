package main

import "fmt"

func main() {
	PrintLog(5 ,"Ali" , false , 78 , "Ehsan" , "Reza" , 8 , true)
}
func PrintLog(logs ...interface{}) {
	for i, item := range logs {
		fmt.Printf("%d : %v \n ",i , item )
	}
}