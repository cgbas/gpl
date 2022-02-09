/*

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
