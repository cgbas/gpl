// HelloWorld exibe uma mensagem de boas-vindas
// aceita argumento ou então usa Neo como nome.
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	var s string
	if len(os.Args) > 1 {
		s = strings.Join(os.Args[1:], " ")
	} else {
		s = "Neo"
	}
	fmt.Printf("%s...\n", s)
	fmt.Printf("Wake up, %s\n", s)
	fmt.Println("The Matrix has you...")
	fmt.Println("Follow the white rabbit.")
}
