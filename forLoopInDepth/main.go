package main

import "fmt"

func main() {
	x := 1
	for x < 10 {
		fmt.Println("x")
		x++
	}
	fmt.Println("Done.")

	// Example of using break
	y := 0
	for {
		y++
		if y > 100 {
			break
		}
		if y%2 != 0 {
			continue
		}
		fmt.Println(y)

	}
	fmt.Println("Done.")

	// How to get the remainder of a
	// % = modulus
	z := 43 % 40
	fmt.Println("z: ", z)

	miniExercise()
	boolLoop()
	initializationStatement()
}

func miniExercise() {
	for r := 33; r <= 122; r++ {
		fmt.Printf("%v\t%#U\n", r, r)
	}
}

func boolLoop() {
	if true {
		fmt.Println("001")
	}
	if false {
		fmt.Println("002")
	}
	if !true {
		fmt.Println("003")
	}
	if !false {
		fmt.Println("004")
	}
}

func initializationStatement() {
	if e := 24; e == 2 {
		fmt.Println("emz")
	}
	fmt.Println("here is a statement")
	fmt.Println("something else")
}
