package main

import "fmt"

func main() {
	x := 42
	y := "James Bond"
	z := true
	s := fmt.Sprintf("%d\t%s\t%t\n", x, y, z)
	fmt.Print(s)

	fmt.Println(x)
	fmt.Printf("My name is Bond, %s\n", y)
	fmt.Println(z)
}
