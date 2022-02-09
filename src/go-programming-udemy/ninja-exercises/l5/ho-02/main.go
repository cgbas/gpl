/*
Take the code from the previous exercise, then store the values of type person in a map with the
key of last name. Access each value in the map. Print out the values, ranging over the slice.
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

	mp := make(map[string]person)

	mp[p1.last] = p1
	mp[p2.last] = p2

	//fmt.Println(mp)
	for _, v := range mp {
		fmt.Println("person", v.last)
		for _, vi := range v.fav {
			fmt.Printf("\t%v\n", vi)
		}
		
	}
}
