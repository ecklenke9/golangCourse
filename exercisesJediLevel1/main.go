package main

import (
	"fmt"
)

// Exercise #2 and #3
// Compiler assigned values for the variables are called the "Zero Value"
var x int = 42
var y string = "James Bond"
var z bool = true

// Creating type emma
type emma int

// Creating a variable of type emma
var e emma
var f int

func main() {
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)

	// Exercise #1
	x := 42
	y := "James Bond"
	z := true

	fmt.Println("X: ", x)
	fmt.Println("Y: ", y)
	fmt.Println("Z: ", z)
	fmt.Printf("X: %v, Y: %v Z: %v \n", x, y, z)

	// Exercise #3
	s := fmt.Sprintf("%v %v %v", x, y, z)
	fmt.Println(s)

	// Exercise #4
	fmt.Println(e)
	fmt.Printf("%T\n", e)
	e = 42
	fmt.Println(e)

	// Exercise #5
	f = int(e)
	fmt.Println(f)
	fmt.Printf("%T\n", f)
}
