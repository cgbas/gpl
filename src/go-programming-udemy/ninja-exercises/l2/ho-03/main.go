// Create TYPED and UNTYPED constants. Print the values of the constants.
package main

import "fmt"

const (
	_     = iota
	a     = iota
	b     = iota
	c     = iota
	d int = iota
	e int = iota
	f int = iota
)

func main() {
	fmt.Println(a, b, c)
	fmt.Printf("%T %T %T\n", a, b, c)
	fmt.Println(d, e, f)
	fmt.Printf("%T %T %T\n", d, e, f)
}
