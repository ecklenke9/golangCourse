package main

import "fmt"

var y = 42

// DECLARE the VARIABLE with the IDENTIFIER "z"
// is of TYPE string
// and I ASSIGN the VALUE "shaken, not stirred"

var z string = "Shaken, not stirred"
var a string = `James said, "Shaken, not stirred"`

// this is a STATIC	programming language
// a VARIABLE is DECLARED to hold a VALUE of a certain TYPE
// not a DYNAMIC programming language

// PRIMITIVE: String, Numeric Type, Bool, and Error
// COMPOSITE: Array, Struct, Pointer, Function, Interface, Slice, Map, and Channel

func main() {
	fmt.Println(y)
	fmt.Println(z)
	fmt.Println(a)
}
