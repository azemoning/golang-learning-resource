package main

import "fmt"

func main() {

	// string type must be inside "" (quotation mark)
	// "hello", is string
	// hello, is not string

	fmt.Println("Hello, again")

	// count total string characters
	fmt.Println("total: ", len("Hello, again"))

	// get single character (will returned with byte not character itself) from string index
	// index start from 0
	fmt.Println("Get 'e': ", "Hello, again"[1])
}
