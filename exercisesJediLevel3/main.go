package main

import "fmt"

func main() {
	//exercise1()
	//exercise2()
	//exercise3()
	//exercise4()
	//exercise5()
	//exercises6and7()
	//exercise8()
	exercise9()
}

func exercise1() {
	for i := 1; i <= 10000; i++ {
		fmt.Println(i)
	}
}

func exercise2() {
	for e := 65; e <= 90; e++ {
		fmt.Println(e)
		for z := 1; z <= 3; z++ {
			fmt.Printf("\t %#U \n", e)
		}
		fmt.Printf("\n")
	}
}

func exercise3() {
	y := 1996
	for y <= 2021 {
		fmt.Printf("%v \n", y)
		y++
	}
}

func exercise4() {
	y2 := 1996
	for {
		fmt.Println(y2)
		y2++
		if y2 > 2021 {
			break
		}
	}
}

func exercise5() {
	for m := 10; m <= 100; m++ {
		fmt.Printf("%v \t %v \n", m, m%4)
	}
}

func exercises6and7() {
	yo := "What's up"
	if yo == "What's up" {
		fmt.Println("Nothing much dawg")
	} else if yo == "Ayo" {
		fmt.Println("It's a good dayoooo")
	} else {
		fmt.Println("Lame.")
	}
}

func exercise8() {
	g := "groot"
	h := "hey"
	switch {

	case g != "groot":
		fmt.Println("GROOOOOOT!!!")
	case h != "hi":
		fmt.Println("Gotcha")
	default:
		fmt.Println("The end is near")
	}
}

func exercise9() {
	favSport := "softball"
	switch favSport {
	case "volleyball":
		fmt.Println("Nah man")
	case "softball":
		fmt.Println("Hell yeah")
	case "football":
		fmt.Println("What the what")
	default:
		fmt.Println("SPORTSSS")
	}
}
