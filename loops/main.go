package main

import "fmt"

func main() {

	//for init; condition; post {}

	// For Loop Example:
	for x := 0; x <= 100; x++ {
		fmt.Println(x)
	}

	// Nested For Loop Example:
	for i := 0; i <= 10; i++ {
		fmt.Printf("The outer loop: %d \n", i)
		for j := 0; j <= 3; j++ {
			fmt.Printf("\t \t The inner loop: %d \n", j)
		}
	}
}
