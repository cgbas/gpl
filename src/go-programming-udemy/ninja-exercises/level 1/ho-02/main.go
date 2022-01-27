package main

import "fmt"

var x int
var y string
var z bool

func main() {
	fmt.Printf("%T\t%T\t%T\n", x, y, z)
	fmt.Printf("'%d'\t'%s'\t'%t'\n", x, y, z)
}
