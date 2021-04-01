package main

import "fmt"

func main() {
	// A missing switch expression is equivalent to the boolean value true
	switch {
	case false:
		fmt.Println("false")
	case (2 == 4):
		fmt.Println("false")
	case (3 == 3):
		fmt.Println("true")
		fallthrough
	// Only falls through the next case. The next case evaluates to true even if it is a false statement.
	case (4 == 1):
		fmt.Println("false, does it print?")
		fallthrough
	case (9 == 9):
		fmt.Println("true, should print")
		fallthrough
	default:
		fmt.Println("this is the default case")
	}

	defaultCase()
}

func defaultCase() {
	switch "Bond" {
	// case followed by expression list
	case "Yo", "Bond":
		// statement list
		fmt.Println("Yo, Bond")
	case "James":
		fmt.Println("James")
	default:
		fmt.Println("Default")
	}
}
