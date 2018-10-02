package main

import (
	"fmt"
)

func main() {
	a := 5
	b := &a
	fmt.Printf("%T -> %T\n", a, b)
	fmt.Println(a, *b)

	// Change value by address
	*b = 10
	fmt.Println(a, *b)
}
