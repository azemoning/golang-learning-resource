package main

import "fmt"

func main() {

	// math operations in golang

	a, b := 20, 10

	// addition
	fmt.Println("A + B: ", a+b)

	// subtraction
	fmt.Println("A - B: ", a-b)

	// division
	fmt.Println("A / B: ", a/b)

	// multiplication
	fmt.Println("A * B: ", a*b)

	// augmented assignment

	var x = 10

	// addition assignment
	x += 10 // equals x = x + 10

	// subtraction assignment
	x -= 10 // equals x = x - 10

	// multiplication assignment
	x *= 10 // equals x = x * 10

	// division
	x /= 10 // equals x = x / 10

	// unary operator

	var z = 10

	// addition
	z++ // equals z = z + 1

	// subtraction
	z-- // equals z = z - 1

}
