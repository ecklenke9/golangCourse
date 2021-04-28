package main

import "fmt"

func main() {
	// This is an entire type
	// It is a composite literal
	// The map contains a [key] and a value
	m := map[string]int{
		"James": 32,
		"Emma":  24,
		"Lora":  29,
	}

	fmt.Println(m)
	fmt.Println(m["James"])

	// If a key that is not stored within the map is entered, it will return the 0 value.
	fmt.Println(m["Spencer"])

	// Comma ok idiom
	v, ok := m["Spencer"]
	fmt.Println(v)
	fmt.Println(ok)

	if v, ok := m["Emma"]; ok {
		fmt.Println("This is the IF print statement", v)
	}

	m["Todd"] = 33

	for k, v := range m {
		fmt.Printf("Key: %v, Value: %v \n", k, v)
	}

	// Slice of int
	xi := []int{1, 3, 5, 7, 9}
	for i, v := range xi {
		fmt.Printf("Index: %v, Value: %v \n", i, v)
	}
	deletingAnElement()
}

func deletingAnElement() {
	m := map[string]int{
		"James": 32,
		"Emma":  24,
		"Lora":  29,
	}

	// Deleting an element from a map:
	// delete(<map name>, "key")
	delete(m, "James")
	fmt.Println(m)

	if v, ok := m["Emma"]; ok {
		fmt.Println("Value: ", v)
		delete(m, "Emma")
	}
	fmt.Println(m)
}
