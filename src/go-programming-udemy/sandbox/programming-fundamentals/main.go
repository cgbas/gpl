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
	_ = iota
	// 2^10, 2^20, 2^30
	// Using iotas for some bit shifting
	kb = 1 << (iota * 10)
	mb = 1 << (iota * 10)
	gb = 1 << (iota * 10)
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

	fmt.Printf("%T\t%T\t%T\t[%T - %d]\t[%T - %d]\n", a, b, c, d, d, e, e)

	fmt.Printf("%T\t(%b)\n", kb, kb)
	fmt.Printf("%T\t(%b)\n", mb, mb)
	fmt.Printf("%T\t(%b)\n", gb, gb)
}
