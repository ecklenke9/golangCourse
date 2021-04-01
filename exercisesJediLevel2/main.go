package main

import (
	"fmt"
)

type hello string

const greeting = hello("what's up")

const (
	ayo = "it's a new dayo"

	year2018 = 2017 + iota
	year2019 = 2017 + iota
	year2020 = 2017 + iota
	year2021 = 2017 + iota
)

func main() {
	// Exercise #1
	number := 42
	// decimal
	fmt.Printf("%d\n", number)
	// binary
	fmt.Printf("%b\n", number)
	// hex
	fmt.Printf("%#x\n", number)

	// Exercise 2
	a := 3 == 3
	b := 6 <= 9
	c := 4 >= 0
	d := 10 != 0
	e := 7 < 9
	f := 21 > 9

	fmt.Println("a: ", a)
	fmt.Println("b: ", b)
	fmt.Println("c: ", c)
	fmt.Println("d: ", d)
	fmt.Println("e: ", e)
	fmt.Println("f: ", f)

	// Exercise 3
	fmt.Println(greeting)
	fmt.Println(ayo)

	// Exercise 4
	bitShift := 4
	fmt.Printf("%d\n", bitShift)
	fmt.Printf("%b\n", bitShift)
	fmt.Printf("%#x\n\n", bitShift)

	// Shifts over 1 place to the left
	nextShift := bitShift << 1
	fmt.Printf("%d\n", nextShift)
	fmt.Printf("%b\n", nextShift)
	fmt.Printf("%#x\n", nextShift)

	// Exercise 5
	// Create a literal of type string using a raw string literal.
	// Print it.
	raw := `here is something as a raw string literal: "something"`
	fmt.Println(raw)

	// Exercise 6
	fmt.Printf("Year 2018: %v \n", year2018)
	fmt.Printf("Year 2019, %v \n", year2019)
	fmt.Printf("Year 2020: %v \n", year2020)
	fmt.Printf("Year 2021: %v \n", year2021)
}
