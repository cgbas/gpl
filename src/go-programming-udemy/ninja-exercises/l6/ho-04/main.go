/*
Create a user defined struct with
	the identifier “person”
	the fields:
		first
		last
		age
	ATTACH a METHOD to TYPE person with the IDENTIFIER “speak”
	the method should
		have the PERSON say their
			NAME and AGE
	create a VALUE of TYPE person
	call the METHOD from the value of type person
*/
package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func main() {
	p := person{
		first: "Michael",
		last:  "Jordan",
		age:   33,
	}
	p.speak()
}

func (s person) speak() {
	fmt.Println("My name is", s.first, s.last, ", and I'm", s.age)
}
