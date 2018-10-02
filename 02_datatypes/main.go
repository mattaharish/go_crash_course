package main

import (
	"fmt"
	"strconv"
)

var name = "Harish Matta"
var age = 23

func sum(y int, z int) int {
	return y + z
}

func statement(name string, age int, isCool bool, email string) string {
	x := "Hi! My name is " + name + "aged " + strconv.Itoa(age) + " and I feel like, I'm cool " + strconv.FormatBool(isCool) + "And my email id " + email
	return x
}

func main() {
	isCool := true
	email := "matta@gmail.com"
	fmt.Printf("Type-> %T\n", name)
	fmt.Printf("Type-> %T\n", age)

	fmt.Print("\n", sum(10, 20))

	fmt.Println(statement(name, age, isCool, email))
}
