/*
Write down what these print:
	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(!true)
*/
package main

import "fmt"

func main() {
	fmt.Printf("true && true\t%v\n", true && true)
	fmt.Printf("true && false\t%v\n", true && false)
	fmt.Printf("true || true\t%v\n", true || true)
	fmt.Printf("true || false\t%v\n", true || false)
	fmt.Printf("!true\t\t%v\n", !true)
}
