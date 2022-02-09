/*
	Closure is when we have “enclosed” the scope of a variable in some code block.
	For this hands-on exercise,
		create a func which “encloses” the scope of a variable:
*/
package main

import "fmt"

func main() {
	inc1 := incr()
	inc2 := incr()

	fmt.Println(inc1())
	fmt.Println(inc1())
	fmt.Println(inc1())

	fmt.Println(inc2())
	fmt.Println(inc2())
}
func incr() func() int {
	var x int
	return func() int {
		x++
		return x
	}
}
