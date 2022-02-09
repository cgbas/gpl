/*
	Build and use an anonymous func
*/
package main

import "fmt"

func main() {
	func(x int) {
		fmt.Println("A resposta é: ", x)
	}(42)
}
