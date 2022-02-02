// Building on the previous hands-on exercise, create a program that uses “else if” and “else”.

package main

import "fmt"

func main() {
	x := "Ludovico Einaudi"
	if x == "Daddy Yankee" {
		fmt.Print("Oye Mami!")
	} else if x == "Ludovico Einaudi" {
		fmt.Println(x)
	} else {
		fmt.Println("Nuthin...")
	}
}
