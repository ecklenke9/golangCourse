package main

import "fmt"

func main() {
	// passing 8 and 9 into addition() as arguments to satisfy
	// the function's parameters
	fmt.Println(addition(8, 9))
}

// function declaration with two parameters (x, y) and a return value (z)
func addition(x int, y int) (z int) {
	z = x + y
	return z
}
