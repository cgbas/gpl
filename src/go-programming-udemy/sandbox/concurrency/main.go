package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	x(1)
}

func x(v int) {
	fmt.Println(v)
	fmt.Println("\t\t", runtime.GOOS)
	fmt.Println("\t\t", runtime.GOARCH)
	fmt.Println("\t\t", runtime.NumCPU())
	fmt.Println("\t\t", runtime.NumGoroutine())
	wg.Add(1)
	go foo()
	wg.Wait()

}

func foo() {
	for i := 0; i <= 10; i++ {
		fmt.Println("bar:", i)
	}
	wg.Done()
}
