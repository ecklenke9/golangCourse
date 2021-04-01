package main

import "fmt"

const (
	_  = iota
	kb = 1024
	mb = 1024 * kb
	gb = 1024 * mb
)

func main() {
	x := 4
	fmt.Printf("%d\t\t%b\n", x, x)

	y := x << 1
	fmt.Printf("%d\t\t%b\n", y, y)

	fmt.Println(kb)
	fmt.Println(mb)
	fmt.Println(gb)

	fmt.Printf("%d\t\t\t%b\n", kb, kb)
	fmt.Printf("%d\t\t\t%b\n", mb, mb)
	fmt.Printf("%d\t\t%b\n", gb, gb)
}
