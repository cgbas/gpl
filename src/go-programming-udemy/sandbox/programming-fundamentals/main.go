package main

import "fmt"

const (
	a int     = 42
	b float64 = 42.78
	c         = "James bond"
	d         = iota
	e         = iota
)
const (
	f = iota
)

func main() {
	s := `Wake up, Neo
	...`
	bs := []byte(s)

	r := "Result: "
	u := "R UTF8: "
	for _, v := range bs {
		r += fmt.Sprintf("%v ", v)
		// Printing the UTF8 codepoint
		u += fmt.Sprintf("%#U ", v)
	}
	r += "\n"
	fmt.Println(r)
	fmt.Println(u)
	fmt.Println(s)

	fmt.Printf("%T\t%T\t%T\t[%T - %d]\t[%T - %d]\t[%T - %d]\n", a, b, c, d, d, e, e, f, f)

}
