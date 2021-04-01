package main

import "fmt"

// DECLARE the variable "y"
// ASSIGN the value 43
// DECLARE and ASSIGN (initialize)
var y = 43

// DECLARE there is a VARIABLE with the IDENTIFIER "z"
// and assigns the VARIABLE with the IDENTIFIER	"z" is of TYPE int
// ASSIGNS the ZERO VALUE of TYPE int to "z"
// 0 for all integer types,
// 0.0 for floating point numbers,
// false for booleans,
// "" for strings,
// nil for interfaces, slices, channels, maps, pointers and functions.
var z int

func main() {
	// short declaration operator
	// DECLARE a variable and ASSIGN a VALUE (of a certain TYPE)
	x := 42
	z = 44

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}
