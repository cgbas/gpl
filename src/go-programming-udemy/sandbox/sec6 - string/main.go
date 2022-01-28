package main

import "fmt"

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

}
