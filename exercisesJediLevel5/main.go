package main

import "fmt"

func main() {
	exercise1and2()
	exercise3()
	exercise4()
}

func exercise1and2() {
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

	fmt.Println(p1.firstName, p1.lastName)
	for _, v := range p1.favIceCreams {
		fmt.Println(v)
	}

	fmt.Println(p2.firstName, p2.lastName)
	for _, v2 := range p2.favIceCreams {
		fmt.Println(v2)
	}

	// -- Exercise 2 Begins --
	fmt.Println("\nEXERCISE 2 \n")

	personMap := map[string]person{
		p1.lastName: p1,
		p2.lastName: p2,
	}

	p1Map := personMap["klenke"]
	p2Map := personMap["tyminski"]

	fmt.Println(p1Map.firstName, p1Map.lastName)
	for _, mv := range p1Map.favIceCreams {
		fmt.Println(mv)
	}

	fmt.Println(p2Map.firstName, p2Map.lastName)
	for _, mv2 := range p2Map.favIceCreams {
		fmt.Println(mv2)
	}

	for _, alt := range personMap {
		fmt.Println(alt.firstName)
		fmt.Println(alt.lastName)
		for _, alt2 := range alt.favIceCreams {
			fmt.Println(alt2)
		}
	}
}

func exercise3() {
	type vehicle struct {
		doors int
		color string
	}

	type truck struct {
		vehicle
		fourWheel bool
	}

	type sedan struct {
		vehicle
		luxury bool
	}

	truck1 := truck{
		vehicle: vehicle{
			doors: 4,
			color: "blue",
		},
		fourWheel: true,
	}

	sedan1 := sedan{
		vehicle: vehicle{
			doors: 2,
			color: "red",
		},
		luxury: false,
	}

	fmt.Println(truck1)
	fmt.Println(sedan1)

	fmt.Println(truck1.color)
	fmt.Println(sedan1.doors)

}

func exercise4() {
	dog := struct {
		name   string
		breed  string
		age    int
		colors []string
		noises map[int]string
	}{
		name:   "CeCe",
		breed:  "Beagle",
		age:    2,
		colors: []string{"brown", "black", "white"},
		noises: map[int]string{
			1: "yap",
			2: "yip",
			3: "aaaarooooohhh",
			4: "bork",
		},
	}

	fmt.Println(dog.name, dog.breed, dog.age)

	for _, v1 := range dog.colors {
		fmt.Println(v1)
	}

	for _, v2 := range dog.noises {
		fmt.Println(v2)
	}
}
