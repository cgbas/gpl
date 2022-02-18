/*
Hands-on exercise #7
write a program that
	launches 10 goroutines
	each goroutine adds 10 numbers to a channel
	pull the numbers off the channel and print them
solutions:
	https://play.golang.org/p/R-zqsKS03P
	https://play.golang.org/p/quWnlwzs2z
	https://go.dev/play/p/WqYnBC_CiKn
video: 170
*/
package main

import "fmt"

func main() {
	c := make(chan int)

	for i := 0; i < 10; i++ {
		go func(i int) {
			for j := 0; j < 10; j++ {
				c <- i*10 + j
			}
		}(i)
	}

	for j := 0; j < 100; j++ {
		fmt.Println(j, <-c)
	}
}
