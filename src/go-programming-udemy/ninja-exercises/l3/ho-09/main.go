/*
Create a program that uses a switch statement with the switch expression specified as a
variable of TYPE string with the IDENTIFIER “favSport”.
*/
package main

import "fmt"

func main() {
	favSport := "Skate"
	switch favSport {
	case "Skate":
		fmt.Println("Do a kickflip!")
	default:
		fmt.Println("Esporte errado :)")
	}
}
