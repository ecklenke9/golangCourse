package main

import "fmt"

var y string
var z int

func main() {
	// DECLARE a variable to be of a certain TYPE
	// and then ASSIGN a VALUE of that type to the variable

	fmt.Println("printing the value of y:", y, "ending")
	// Printing out the type of y
	fmt.Printf("%T\n", y)

	y = "Emma, Klenke"

	fmt.Println(y)
	fmt.Printf("%T\n", y)

	fmt.Println(z)
	fmt.Printf("%T\n", z)

}
