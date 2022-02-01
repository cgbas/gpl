// Using iota, create 4 constants for the NEXT 4 years. Print the constant values
package main

import "fmt"

const (
	current = 2022
	a       = (current + iota)
	b       = (current + iota)
	c       = (current + iota)
	d       = (current + iota)
)

func main() {
	fmt.Println(a, b, c, d)
}
