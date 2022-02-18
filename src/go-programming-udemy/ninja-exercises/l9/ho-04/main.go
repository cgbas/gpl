/*
	Fix the race condition you created in the previous exercise by using a mutex
		it makes sense to remove runtime.Gosched()
*/

package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Goroutines:", runtime.NumGoroutine())

	incrementer := 0

	const gs = 100
	var wg sync.WaitGroup
	var m sync.Mutex
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			m.Lock()
			v := incrementer
			// time.Sleep(time.Second)
			v++
			incrementer = v
			m.Unlock()
			wg.Done()
		}()
		fmt.Println("Goroutines:", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	fmt.Println("count:", incrementer)
}
