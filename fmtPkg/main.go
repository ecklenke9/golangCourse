package main

import "fmt"

// %v	the value in a default format
// when printing structs, the plus flag (%+v) adds field names
// %#v	a Go-syntax representation of the value
// %T	a Go-syntax representation of the type of the value
// %%	a literal percent sign; consumes no value

var y = 42

func main() {
	fmt.Println(y)
	fmt.Printf("%T\n", y)
	fmt.Printf("%b\n", y)
	fmt.Printf("%x\n", y)
	fmt.Printf("%#x\n", y)
	fmt.Printf("%#x\t%b\t%x\n", y, y, y)

	// String print
	s := fmt.Sprintf("%#x\t%b\t%x\n", y, y, y)
	fmt.Println(s)
	fmt.Printf("%v\n", y)
}
