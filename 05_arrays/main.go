package main

import (
	"fmt"

	"github.com/mattaharish/go_crash/03_package/numutil"
)

func main() {
	// Arrays
	a := []float32{1, 2, 3, 4, 5, 6, 7, 8}
	var avg float32
	avg = numutil.AvgOfArray(a)
	// fmt.Println(numutil.SumOfArray(a))
	fmt.Println(avg)

}
