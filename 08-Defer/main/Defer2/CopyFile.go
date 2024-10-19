package main

import (
	"io"
	"os"
)

func main() {





	copyFile("D:\\go-project\\08-Defer\\main\\Defer2\\destinationName.txt" ,"D:\\go-project\\08-Defer\\main\\Defer2\\sourceName.txt" )

	
}
func copyFile(destinationName,sourceName string) (written int64, err error) {
	source,err := os.Open(sourceName)
	if err!=nil{
		return
	}
	defer source.Close()
	destination,err := os.Create(destinationName)
	if err!=nil{
		return
	}
	defer destination.Close()
	return io.Copy(destination, source)
}