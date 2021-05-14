package main

import "fmt"

func main() {
	exercise1()
}

func exercise1() {
	type person struct {
		firstName    string
		lastName     string
		favIceCreams []string
	}
	p1 := person{
		firstName:    "emma",
		lastName:     "klenke",
		favIceCreams: []string{"cookie dough", "strawberry", "vanilla"},
	}

	p2 := person{
		firstName:    "spencer",
		lastName:     "tyminski",
		favIceCreams: []string{"chocolate", "rocky road", "mint choc chip"},
	}

	fmt.Println(p1)
	fmt.Println(p2)
}
