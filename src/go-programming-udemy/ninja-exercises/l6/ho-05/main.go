/*
	create a type SQUARE
	create a type CIRCLE
	attach a method to each that calculates AREA and returns it
	circle area= π r 2
	square area = L * W
	create a type SHAPE that defines an interface as anything that has the AREA method
	create a func INFO which takes type shape and then prints the area
	create a value of type square
	create a value of type circle
	use func info to print the area of square
	use func info to print the area of circle
*/
package main

import (
	"fmt"
	"math"
)

type square struct {
	l, w float64
}
type circle struct {
	r float64
}
type shape interface {
	area()
}

func info(s shape) {
	switch s.(type) {
	case circle:
		s.(circle).area()
	case square:
		s.(square).area()
	}

}

func (s circle) area() {
	fmt.Println("Area is: ", math.Pi*(math.Pow(s.r, 2)))
}

func (s square) area() {
	fmt.Println("Area is: ", s.l*s.w)
}

func main() {
	s := square{
		l: 10,
		w: 10,
	}
	c := circle{
		r: 2,
	}
	info(s)
	info(c)
}
