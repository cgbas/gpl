/*

Hands-on exercise #4
Starting with this code, pull the values off the channel using a select statement
https://play.golang.org/p/MvL6uamrJP

solution: https://play.golang.org/p/FulKBY5JNj
video: 167

*/

package main

import (
	"fmt"
)

func main() {
	q := make(chan int)
	c := gen(q)

	receive(c, q)

	fmt.Println("about to exit")
}

func gen(q chan<- int) <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		q <- 1
		close(c)

	}()

	return c
}

func receive(c, q <-chan int) {

	for {
		select {
		case v := <-c:
			fmt.Println("From c", v)
		case <-q:
			return
		}

	}
}
