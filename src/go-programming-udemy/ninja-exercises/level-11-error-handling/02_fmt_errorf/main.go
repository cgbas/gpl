/*
Start with this code: https://go.dev/play/p/9a1IAWy5E6.

Create a custom error message using “fmt.Errorf”.
solution:
	https://play.golang.org/p/HugU4HJEEO
	https://play.golang.org/p/NII-lmGasj
	https://play.golang.org/p/Vo5kIoR-sG

*/

package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type person struct {
	First   string
	Last    string
	Sayings []string
}

func main() {
	p1 := person{
		First:   "James",
		Last:    "Bond",
		Sayings: []string{"Shaken, not stirred", "Any last wishes?", "Never say never"},
	}

	bs, err := toJSON(p1)
	if err != nil {

		log.Fatal("Error converting to JSon")
	}

	fmt.Println(string(bs))

}

// toJSON needs to return an error also
func toJSON(a interface{}) ([]byte, error) {
	bs, err := json.Marshal(a)
	if err != nil {
		return []byte{}, fmt.Errorf("there was an error in to json: %v", err)
	}
	return bs, nil
}
