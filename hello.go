package main

import "fmt"

func main() {
	fmt.Println("Hello this is main")
	food()
	bars()
	trickery()
}
func food() {
	fmt.Println("This is food")
}
func bars() {
	fmt.Println("This is bars")
}

type foo int

var y foo

const bar int = 42
const yogurt = "strawberry"

func trickery() {
	y = 42
	x := yogurt
	fmt.Printf("%T \n", y)
	fmt.Printf("%T \n", int(y))
	fmt.Printf("%T \n", bar)
	fmt.Println(bar)
	fmt.Printf("%T", x)
}

//Control Flow
// 1. Sequence
// 2. Loop; Iterative
// 3. Conditional
