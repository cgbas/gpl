/*
	Use the “defer” keyword to show that a deferred func runs after the func containing it exits.
*/
package main

import "fmt"

func main() {
	defer fmt.Println(f1())
	fmt.Println(f2())
}

func f1() string {
	return "that's f1"
}
func f2() string {
	return "that's f2"
}
