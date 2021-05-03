package main

import "fmt"

func main() {
	exercise1()
	exercise2()
	exercise3()
}

func exercise1() {
	arrayOfInts := [5]int{1, 2, 3, 4, 5}

	for i, v := range arrayOfInts {
		fmt.Printf("Index: %v, Value: %v \n", i, v)
	}

	fmt.Printf("Type: %T \n", arrayOfInts)
}

func exercise2() {
	sliceOfInts := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	for i, v := range sliceOfInts {
		fmt.Printf("Index: %v, Value: %v \n", i, v)
	}

	fmt.Printf("Type: %T \n", sliceOfInts)
}

func exercise3() {
	sliceOfInts := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	fmt.Println(sliceOfInts[:5])
	fmt.Println(sliceOfInts[5:])
	fmt.Println(sliceOfInts[2:7])
	fmt.Println(sliceOfInts[1:6])
}
