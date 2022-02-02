package main

import "fmt"

func main() {
	// array example
	x := []int{0, 0, 0, 0}
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

	for i := 0; i < len(z); i++ {
		fmt.Println(z[i])
	}

	// slicing a slice
	fmt.Println(z[1:3])
	fmt.Println(z)
	z = append(z, x...)
	fmt.Println(z)
}
