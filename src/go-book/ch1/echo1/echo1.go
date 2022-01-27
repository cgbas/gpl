// Echo1 exibe os argumentos da linha de comando
package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	for i, v := range os.Args[1:] {
		s += sep + os.Args[i+1]
		sep = " "
		fmt.Printf("[%v] = %v\n", i+1, v)
	}
	fmt.Println(s)
}
