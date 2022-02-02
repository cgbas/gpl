// Print out the remainder (modulus) which is found for each number between 10 and 100 when it
// is divided by 4.
package main

import "fmt"

func main() {
	i := 10
	for i <= 100 {
		fmt.Printf("%v\t%% 4 =\t%v\n", i, i%4)
		i++
	}

}
