/* Create a for loop using this syntax
	for { }
Have it print out the years you have been alive.
*/
package main

import "fmt"

func main() {
	bd := 1987
	for {
		if bd > 2022 {
			break
		}
		fmt.Println(bd)
		bd++
	}
}
