// FizzBuzz

package main

import (
	"fmt"
)

func main() {
	for index := 1; index <= 15; index++ {
		if index%3 == 0 && index%5 == 0 {
			fmt.Printf("%d -> fizzBuzz\n", index)
		} else if index%5 == 0 {
			fmt.Printf("%d -> Buzz\n", index)
		} else if index%3 == 0 {
			fmt.Printf("%d is Fizz\n", index)
		} else {
			fmt.Printf("XXXX\n")
		}
	}
}
