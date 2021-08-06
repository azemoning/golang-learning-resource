package main

import "fmt"

func main() {
	// we can do type conversion in golang
	// example: int32 to int64

	var value32 int32 = 100000
	var value64 int64 = int64(value32)

	fmt.Println("value with int64: ", value64)

	// becareful not to convert the value that exceeding the type limit

	// example: convert byte from string index to character

	var name = "John"
	var oByte = name[1]
	var oChar = string(oByte)

	fmt.Println("o character from 'John': ", oChar)
}
