/*
	Assign a func to a variable, then call that func
*/
package main

import "fmt"

func main() {
	f := func(s string) {
		fmt.Println("Hello,", s)
	}
	f("Orwell")

}
