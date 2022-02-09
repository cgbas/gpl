/*
Create your own type “person” which will have an underlying type of “struct” so that it can store
the following data:
	first name
	last name
	favorite ice cream flavors
Create two VALUES of TYPE person. Print out the values, ranging over the elements in the slice
which stores the favorite flavors
*/
package main

import "fmt"

type person struct {
	first string
	last  string
	fav   []string
}

func main() {
	p1 := person{"John", "Crawl", []string{"Strawberry", "Vanilla", "Sambayón"}}
	// more idiomatic
	p2 := person{
		first: "Duke",
		last:  "Barnes",
		fav: []string{
			"Chocolate",
			"Vanilla",
			"Coconut",
		},
	}

	pp := []person{p1, p2}
	fmt.Println(p1, p2)
	for _, v := range pp {
		fmt.Println("Flavors for", v.first, v.last)
		for ii, vv := range v.fav {
			fmt.Println("\t", ii, vv)
		}
	}
}
