package main

import "fmt"

func main() {
	// array example
	x := []int{0, 0, 0, 0}
	fmt.Println(len(x))
	var y [6]int
	fmt.Println(len(y))

	// composite literal
	// v := type{values}
	z := []int{4, 5, 7, 8, 42}
	fmt.Println(z)

	// looping over an slice
	for key, value := range z {
		fmt.Println(key, value)
	}

	for i := 0; i < len(z); i++ {
		fmt.Println(z[i])
	}

	// slicing a slice
	fmt.Println(z[1:3])
	fmt.Println(z)
	z = append(z, x...)
	fmt.Println(z)

	z = append(z[:2], z[4:]...)
	fmt.Println(z)

	// playing with make
	a := make([]int, 10, 100)
	a = append(a, 18)
	fmt.Println(a)
	fmt.Println(cap(a))
	fmt.Println(len(a))

	// multidimensional
	ab := [][]int{z, a}
	fmt.Println(ab)

	// maps
	m := map[string]int{
		"James":           32,
		"Miss Moneypenny": 27,
	}
	fmt.Println(m)
	fmt.Println(m["James"])
	fmt.Println(m["Chris"])

	// comma ok idiom
	if v, ok := m["James"]; ok {
		fmt.Println("THIS IS THE IF PRINT", v)
	}
	delete(m, "James")
	fmt.Println(m)

}
