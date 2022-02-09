/*
	Closure is when we have “enclosed” the scope of a variable in some code block.
	For this hands-on exercise,
		create a func which “encloses” the scope of a variable:
*/
package main

import "fmt"

func main() {
	xf := f()
	xfr := xf()
	fmt.Println(xfr)

}
func f() func() string {
	return func() string {
		return "Rawr"
	}
}
