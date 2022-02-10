/*
	create a person struct
		create a func called “changeMe” with a *person as a parameter
		change the value stored at the *person address
	■ important
		to dereference a struct, use (*value).field
		p1.first
		(*p1).first
		“As an exception, if the type of x is a named pointer type and (*x).f
			is a valid selector expression denoting a field (but not a method),
			x.f is shorthand for (*x).f.”
			https://golang.org/ref/spec#Selectors
		create a value of type person
		print out the value
		in func main
		call “changeMe”
		in func main
		print out the value
*/

package main

import "fmt"

type person struct {
	last_first string
	address    string
}

func changeMe(p *person, a string) {
	(*p).address = a
	fmt.Printf("changed to:%v\t%T\n", p.address, p)
}

func main() {
	p := person{
		address: "Highway 1",
	}
	fmt.Println(p.address)
	changeMe(&p, "Highway 2")
	changeMe(&p, "Highway 3")
}
