package main

import "fmt"

// A composite data type is any data type which can be constructed
// in a program using the programming language's primitive data types
// and other composite types.

// Aggregates can be either an array or slice holding values of the same type.

// A SLICE allows you to group together VALUES of the same TYPE.

func main() {
	// Composite Literal
	// x := type{values}

	x := []int{4, 5, 6, 9}
	fmt.Println(x)

	//loopingOverSlice()
	//slicingASlice()
	//appendToASlice()
	//deletingFromASlice()
	//sliceUsingMake()
	multiDimensionalSlice()
}

func loopingOverSlice() {
	y := []int{7, 9, 23, 87}
	fmt.Println(len(y))
	fmt.Println(y)
	fmt.Println(y[3])

	for i, v := range y {
		fmt.Printf("Index: %v, Value: %v \n", i, v)
	}
}

func slicingASlice() {
	e := []int{8, 6, 3, 64, 7}
	fmt.Println(e)
	fmt.Println(e[1])
	// Slicing the slice. Grabs index 1 until end.
	fmt.Println(e[1:])
	// The index at the right side of the colon goes up to but does
	// not include the index.
	fmt.Println(e[1:4])

	for i := 0; i <= 4; i++ {
		fmt.Println(e[i])
	}
}

// func append(slice []T, elements ...T) []T
// where T is a placeholder for any given type

func appendToASlice() {
	a := []int{3, 4, 5, 6, 7, 8}
	fmt.Println(a)
	a = append(a, 9, 10, 11, 12, 13)
	fmt.Println(a)

	d := []int{234, 456, 678, 890}

	// ... is considered a variadic parameter
	a = append(a, d...)
	fmt.Println(a)
}

func deletingFromASlice() {
	r := []int{7, 8, 9, 4, 2, 5}
	fmt.Println(r)

	// Delete all the way up to but not including the index of 2
	r = append(r[:2], r[4:]...)
	fmt.Println(r)
}

func sliceUsingMake() {
	// make(type, length, capacity)
	w := make([]int, 12, 100)
	fmt.Println(w)
	fmt.Println(len(w))
	fmt.Println(cap(w))
	w[0] = 3
	w[9] = 999
	fmt.Println(w)

	w = append(w, 7654)
	fmt.Println(w)
}

func multiDimensionalSlice() {
	jamesBond := []string{"James", "Bond", "chocolate", "martini"}
	fmt.Println(jamesBond)

	moneyPenny := []string{"Miss", "Moneypenny", "strawberry", "bubblegum"}
	fmt.Println(moneyPenny)

	xp := [][]string{jamesBond, moneyPenny}
	fmt.Println(xp)
}
