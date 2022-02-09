/*

 */
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
