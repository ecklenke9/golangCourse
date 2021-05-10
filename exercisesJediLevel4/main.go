package main

import "fmt"

func main() {
	exercise1()
	exercise2()
	exercise3()
	exercise4()
	exercise5()
	exercise6()
	exercise7()
	exercises8and9and10()
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

func exercise4() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println(x)

	x = append(x, 52)
	fmt.Println(x)

	x = append(x, 53, 54, 55)
	fmt.Println(x)

	y := []int{56, 57, 58, 59, 60}
	x = append(x, y...)
	fmt.Println(x)
}

func exercise5() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println(x)

	x = append(x[:3], x[6:]...)
	fmt.Println(x)
}

func exercise6() {
	states := make([]string, 50, 50)
	states = []string{"Alabama", "Alaska",
		"Arizona", "Arkansas", "California", "Colorado", "Connecticut", "Delaware",
		"Florida", "Georgia", "Guam", "Hawaii", "Idaho", "Illinois",
		"Indiana", "Iowa", "Kansas", "Kentucky", "Louisiana", "Maine", "Maryland",
		"Massachusetts", "Michigan", "Minnesota", "Mississippi",
		"Missouri", "Montana", "Nebraska", "Nevada", "New Hampshire", "New Jersey", "New Mexico",
		"New York", "North Carolina", "North Dakota", "Ohio",
		"Oklahoma", "Oregon", "Pennsylvania", "Rhode Island", "South Carolina",
		"South Dakota", "Tennessee", "Texas", "Utah", "Vermont",
		"Virginia", "Washington", "West Virginia", "Wisconsin", "Wyoming"}
	fmt.Println(states)

	for i := 0; i < len(states); i++ {
		fmt.Printf("Index: %v, State: %v \n", i, states[i])
	}
}

func exercise7() {
	jamesBond := []string{"James", "Bond", "Shaken, not stirred"}
	missMoney := []string{"Miss", "MoneyPenny", "Helloooo, James"}
	x := [][]string{jamesBond, missMoney}

	for i, v := range x {
		fmt.Printf("Record: %v \n", i)
		for i2, v2 := range v {
			fmt.Printf("\t Index Position: %v \t Value: %v \n", i2, v2)
		}
	}
}

func exercises8and9and10() {
	m := map[string][]string{
		"bond_james":      {"shaken, not stirred", "martinis", "women"},
		"moneypenny_miss": {"james bond", "literature", "computer science"},
		"no_dr":           {"being evil", "ice cream", "sunsets"},
	}

	for i, v := range m {
		fmt.Printf("This is the record for %v \n", i)
		for i2, v2 := range v {
			fmt.Printf("\t Index: %v \t Value: %v \n", i2, v2)
		}
	}

	m["klenke_emma"] = []string{"coffee", "late nights", "dogs"}
	fmt.Println("\nAdding Element: ")
	for i, v := range m {
		fmt.Printf("This is the record for %v \n", i)
		for i2, v2 := range v {
			fmt.Printf("\t Index: %v \t Value: %v \n", i2, v2)
		}
	}

	delete(m, "no_dr")
	fmt.Println("\nDeleting Element: ")
	for i, v := range m {
		fmt.Printf("This is the record for %v \n", i)
		for i2, v2 := range v {
			fmt.Printf("\t Index: %v \t Value: %v \n", i2, v2)
		}
	}
}
