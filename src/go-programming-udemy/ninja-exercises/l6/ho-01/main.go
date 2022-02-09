/*
Hands on exercise
	create a func with the identifier foo that returns an int
	create a func with the identifier bar that returns an int and a string
	call both funcs
	print out their results
*/

package main

import "fmt"

func main() {

	fmt.Println(foo(1))
	fmt.Println(bar(2, "Mike"))
}

func foo(i int) int {
	return i
}

func bar(i int, s string) (int, string) {
	return i, s
}
