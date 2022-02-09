/*
create a func with the identifier foo that
	takes in a variadic parameter of type int
	pass in a value of type []int into your func (unfurl the []int)
	returns the sum of all values of type int passed in
create a func with the identifier bar that
	takes in a parameter of type []int
	returns the sum of all values of type int passed in
*/

package main

import "fmt"

func main() {

	x := []int{5, 6, 7, 8, 9}
	fmt.Println("foo", foo(x...))
	fmt.Println("bar", bar(x))
}

func foo(p ...int) int {
	sum := 0
	for _, v := range p {
		sum += v
	}

	return sum
}

func bar(p []int) int {
	sum := 0
	for _, v := range p {
		sum += v
	}

	return sum
}
