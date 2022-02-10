/*
	in addition to the main goroutine, launch two additional goroutines
	each additional goroutine should print something out
	use waitgroups to make sure each goroutine finishes before your program exists
*/

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type person struct {
	first string
	last  string
	age   int
}

type quote struct {
	author  string
	quote   string
	details string
}

var wg sync.WaitGroup

func main() {
	p := person{
		first: "Charles",
		last:  "Bukowski",
		age:   30,
	}

	q := []quote{
		{`Tenessee Willians`, `Life is a Journey, not a Destination.`, `Appears on movie Transiberian`},
		{`Henry David Thoreau`, `More Than Love, Than Money, Than Fame, Give Me Truth`, `Appears on movie Into the Wild`},
		{`Alan Moore`, `Behind this mask there's more than flesh. There is an Idea. And Ideas are bulletproof.`, `V's speech`},
	}

	wg.Add(2)

	go p.quote(q...)
	go p.speak()

	wg.Wait()
}

func (p person) speak() {
	fmt.Println("It's", p.first, p.last)
	wg.Done()
}

func (p person) quote(quotes ...quote) {
	rand.Seed(time.Now().UnixNano())
	q := quotes[rand.Intn(len(quotes))]
	/*
		qx := &q
		append(*qx, quote{
			author:  "J. R. R. Tolkien",
			quote:   "Not all those who wander are lost.",
			details: "I talks about Aragorn and it's lineage",
		})
	*/
	fmt.Printf("Author: %s\n\t\"%s\"\n", q.author, q.quote)
	wg.Done()
}
