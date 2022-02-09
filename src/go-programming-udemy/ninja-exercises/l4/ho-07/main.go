/*
	Create a slice of a slice of string ([][]string). Store the following data in the multi-dimensional
	slice:
		"James", "Bond", "Shaken, not stirred"
		"Miss", "Moneypenny", "Helloooooo, James."
	Range over the records, then range over the data in each record.
*/
package main

import "fmt"

func main() {

	x1 := []string{"James", "Bond", "Shaken, not stirred"}
	x2 := []string{"Miss", "Moneypenny", "Helloooooo, James."}
	z := [][]string{x1, x2}

	for i, v := range z {
		fmt.Println(i, v)
		for i, v := range v {
			fmt.Println("\t", i, v)
		}

	}

}
