package main

import "fmt"

func main() {
	s := "Hello"
	t := `Hello
	there`

	fmt.Println(s)
	fmt.Println(t)
	fmt.Printf("%T\n", s)
	fmt.Printf("%T\n", t)

	// Converting var s to a byte slice
	bs := []byte(s)
	fmt.Println(bs)
	fmt.Printf("%T\n", bs)

	// Assigning the 0 index value of s to bbs
	bbs := []byte(s)[0]
	fmt.Println(bbs)

	for i := 0; i < len(s); i++ {
		fmt.Printf("%#U \n", s[i])
	}

	for i, v := range s {
		fmt.Printf("At index position %d we have hex %#x \n", i, v)
	}

}
