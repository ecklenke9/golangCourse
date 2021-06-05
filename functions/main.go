package main

import "fmt"

func main() {
	// We call our func and pass in arguments (if any).
	foo()
	bar("Emma")
	s1 := woo("Lulu")
	fmt.Println(s1)
	x, y := mouse("Mary", "Kay")
	fmt.Println(x)
	fmt.Println(y)
}

// func(r receiver) identifier(parameters) (return(s)) { ... }
// Functions are defined with parameters (if any).
func foo() {
	fmt.Println("hello from foo")
}

// EVERYTHING in Go is passed by VALUE
func bar(s string) {
	fmt.Printf("Hello, %v \n", s)
}

func woo(s string) string {
	return fmt.Sprint("Hello from ", s)
}

func mouse(fn string, ln string) (string, bool) {
	a := fmt.Sprint(fn, " ", ln, ` says, "Hello!"`)
	b := false
	return a, b
}
