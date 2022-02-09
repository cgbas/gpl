/*
	Create and use an anonymous struct
*/
package main

import "fmt"

type sponsor struct {
	start       int
	end         int
	fixedIncome float64
}

func main() {
	as := struct {
		last_first string
		truck_size int
		shape_size float64
		sponsors   map[string]sponsor
	}{
		last_first: "lemos_tiago",
		truck_size: 139,
		shape_size: 7.9,
		sponsors: map[string]sponsor{
			"New Balance": sponsor{
				start:       2010,
				fixedIncome: 500,
			},
			"Jah griptape": sponsor{
				start:       2021,
				fixedIncome: 100,
			},
		},
	}

	fmt.Println("Skater\t", as.last_first)
	fmt.Println("Truck\t", as.truck_size)
	fmt.Println("Shape\t", as.shape_size)
	fmt.Println("Sponsors\t", as.sponsors)

}
