package main

import "fmt"

func main() {
	// array example
	var x [5]int
	fmt.Println(len(x))
	var y [6]int
	fmt.Println(len(y))

	// composite literal
	// v := type{values}
	z := []int{4, 5, 7, 8, 42}
	fmt.Println(z)

	// looping over an slice
	for key, value := range z {
		fmt.Println(key, value)
	}
}
