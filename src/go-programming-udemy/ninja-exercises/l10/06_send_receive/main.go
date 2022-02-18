/*

Hands-on exercise #6
● write a program that
○ puts 100 numbers to a channel
○ pull the numbers off the channel and print them
solution: https://play.golang.org/p/096Lk1BR7o
video: 169

*/

package main

import "fmt"

func main() {
	c := gen()
	receive(c)
}

func receive(c <-chan int) {
	for v := range c {
		fmt.Println(v)
	}
}
func gen() <-chan int {
	c := make(chan int)
	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
	}()
	return c
}
