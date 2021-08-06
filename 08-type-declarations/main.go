package main

import "fmt"

func main() {
	// type declaration enable us to make an alias to existing type
	// example: aliasing a string type

	type UserID string

	var userId UserID = "asdjkl123"
	fmt.Println("user id: ", userId)
}
