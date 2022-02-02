// Create a program that uses a switch statement with no switch expression specified.
package main

import "fmt"

func main() {
	switch {
	case true:
		fmt.Println("Should print")
		fallthrough
	default:
		fmt.Println("Empty switch's default clause")
	}
}
