/*
	unmarshal the JSON into a Go data structure.
*/

package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	First string `json:"First"`
	Age   int    `json:"Age"`
}

func main() {
	s := `[{"First":"James","Age":32},{"First":"Moneypenny","Age":27},{"First":"M","Age":54}]`
	bs := []byte(s)
	fmt.Printf("%T\n", bs)

	var uu = []user{}
	err := json.Unmarshal(bs, &uu)
	if err != nil {
		fmt.Println("Unmarshal err:", err)
	}
	fmt.Println(uu)
}
