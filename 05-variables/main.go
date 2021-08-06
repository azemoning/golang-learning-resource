package main

import "fmt"

func main() {

	// there is some ways to declare variable

	// declaring variable with its type and value initialization
	var firstName string = "John"
	fmt.Println("first name: ", firstName)

	// declaring variable with only its type
	// if we print, it will print out empty string
	var lastName string
	fmt.Println("last name: ", lastName)

	// declaring variable with only value initialization
	// go will set the variable type with the value we entered
	var city = "New York"
	fmt.Println("city: ", city)

	var age = 25
	fmt.Println("age: ", age)

	// declaring variable without keyword "var"
	// its the same like we declaring variable without initiate the type
	// if we do this way, we need to add ":" (colon) before "="
	address := "Highest building in New York"
	fmt.Println("address: ", address)

	phoneNumber := 90123456
	fmt.Println("phone number: ", phoneNumber)

	// declaring multiple variables
	var (
		email    = "john@example.com"
		facebook = "john doe"
		twitter  = "@johndoe"
	)

	fmt.Println("email address: ", email)
	fmt.Println("facebook name: ", facebook)
	fmt.Println("twitter username:", twitter)
}
