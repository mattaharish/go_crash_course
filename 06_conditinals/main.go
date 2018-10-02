package main

import (
	"fmt"
)

func main() {
	x := 10
	y := 15
	if x < y {
		fmt.Printf("%d is less than %d \n", x, y)
	} else {
		fmt.Printf("%d is greater than %d \n", x, y)
	}

	color := "blue"

	switch color {
	case "red":
		fmt.Printf("Your choice is %s", color)
	case "blue":
		fmt.Printf("Your choice is %s \n", color)
	default:
		fmt.Println("Neither...")
	}
}
