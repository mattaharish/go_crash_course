package main

import (
	"fmt"
	"math"

	"github.com/mattaharish/go_crash/03_package/numutil"
)

func squre(n int) float64 {
	return math.Sqrt(25)
}

func main() {
	n := 10
	fmt.Printf("%T", n)
	fmt.Println(squre(n))
	fmt.Println(numutil.Prime(11))
	fmt.Println(numutil.Odd(10))
	fmt.Println(numutil.Even(10))
}
