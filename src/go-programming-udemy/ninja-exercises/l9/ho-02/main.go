/*
	This exercise will reinforce our understanding of method sets:
		create a type person struct
		attach a method speak to type person using a pointer receiver
			*person
		create a type human interface
		to implicitly implement the interface, a human must have the speak method
		create func “saySomething”
		have it take in a human as a parameter
		have it call the speak method
		show the following in your code
		you CAN pass a value of type *person into saySomething
		you CANNOT pass a value of type person into saySomething
		here is a hint if you need some help
		https://play.golang.org/p/FAwcQbNtMG
		Receivers Values
		-----------------------------------------------
		(t T) T and *T
		(t *T) *T
*/

package main

import "fmt"

type person struct {
	name string
}

type human interface {
	speak()
}

func (p *person) speak() {
	fmt.Printf("Hello my name is %s\n", p.name)
}

func saySomething(h human) {
	h.speak()
}

func main() {
	p := person{
		name: "Christian Silva",
	}

	saySomething(&p)
	// Doesn't work
	//saySomething(p)
}
