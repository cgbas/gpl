/*
	A “callback” is when we pass a func into a func as an argument.
	For this exercise,
		pass a func into a func as an argument
*/
package main

import "fmt"

func main() {
	evenSum(sum, 1, 2, 3, 4, 5, 6)
}

func evenSum(f func(x ...int) int, vi ...int) {
	var y []int
	for _, v := range vi {
		if v%2 == 0 {
			y = append(y, v)
		}
	}
	fmt.Println(f(y...))
}

func sum(x ...int) int {
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}
