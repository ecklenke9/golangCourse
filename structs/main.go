package main

import "fmt"

// A struct is a data structure that allows for composing together
// data of different types. It is an aggregate data type. It aggregates together
// values of different types.

// A struct is a sequence of named elements, called fields, each of which has a name and a
// type. Within a struct, non-blank field names must be unique.

type person struct {
	//identifier, type
	first string
	last  string
	age   int
}

type secretAgent struct {
	//person gets promoted to the outer type secretAgent
	person
	ltk bool
}

func main() {
	//structIntro()
	//embeddedStructs()
	anonymousStructs()
}

func structIntro() {
	// Creating a value of type person and assigning the value to p1
	p1 := person{
		first: "James",
		last:  "Bond",
		age:   32,
	}

	// Creating a value of type person and assigning the value to p2
	// A composite literal
	p2 := person{
		first: "Emma",
		last:  "Klenke",
		age:   24,
	}

	fmt.Println(p1)
	fmt.Println(p2)

	fmt.Println(p1.first, p1.last)
	fmt.Println(p2.first, p2.last)
}

func embeddedStructs() {
	sa1 := secretAgent{
		person: person{
			first: "Spencer",
			last:  "Tyminski",
			age:   30,
		},
		ltk: true,
	}

	fmt.Println(sa1)
	fmt.Printf("SA First: %v \nSA Last: %v \nSA Age: %v \nLTK: %v", sa1.first, sa1.last, sa1.age, sa1.ltk)
}

func anonymousStructs() {
	// composite literal
	p1 := struct {
		first string
		last  string
		age   int
	}{
		// values
		first: "Cece",
		last:  "The Beagle",
		age:   3,
	}

	fmt.Println(p1)
}
