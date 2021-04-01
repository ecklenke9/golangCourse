package main

import "fmt"

// A boolean type represents the set of Boolean true values
// denoted by the predeclared constants true and false.
var x bool

func main() {
	fmt.Println(x)
	x = true
	fmt.Println(x)

	a := 7
	b := 42
	fmt.Println(a == b)

}
