package main

import (
	"fmt"
)

func main() {
	// Maps
	userDetails := make(map[int]string)

	userDetails[155] = "harish@flowace.in"
	userDetails[156] = "matta@flowace.in"
	userDetails[157] = "kumar@flowace.in"

	fmt.Println(userDetails)

	for index, data := range userDetails {
		fmt.Printf("%d -> %s\n", index, data)
	}
}
