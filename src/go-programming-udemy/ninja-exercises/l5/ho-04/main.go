/*
	Create and use an anonymous struct
*/
package main

import "fmt"

func main() {
	as := struct {
		last_first string
		truck_size int
		shape_size float64
	}{
		last_first: "lemos_tiago",
		truck_size: 139,
		shape_size: 7.9,
	}

	fmt.Println("Skater\t", as.last_first)
	fmt.Println("Truck\t", as.truck_size)
	fmt.Println("Shape\t", as.shape_size)

}
