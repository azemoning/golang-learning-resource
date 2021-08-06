package main

func main() {
	// constant is a variable that cannot be change after initialization (immutable).
	// ways to declaring a constant
	const firstName string = "John"
	const lastName = "Doe"
	const age = 20
	const phi = 3.14

	// if we declare constant and not used it, golang will not complaint to us

	// declare multiple constants
	const (
		BLACK_COLOR  = "black"
		YELLOW_COLOR = "yellow"
		RED_COLOR    = "red"
	)
}
