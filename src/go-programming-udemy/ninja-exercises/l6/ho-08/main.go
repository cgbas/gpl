/*
	Create a func which returns a func
		assign the returned func to a variable
		call the returned func
*/
package main

import "fmt"

func main() {
	// calling a function returned by another function
	fmt.Println(f()())
}
func f() func() string {
	return func() string {
		return "rawr"
	}
}
