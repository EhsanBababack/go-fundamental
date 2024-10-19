package main

import "time"

func main() {
	names := []string{"Ali", "Saeid", "Reza", "Ehsan", "Sara", "Nickan", "Zahra", "Mona"}

	for i, item := range names {
		go func(index int, name string) {
			println(index, name)
		}(i, item)
	}
	time.Sleep(time.Second*1)
}