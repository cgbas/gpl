/*

Create a struct “customErr” which implements the builtin.error interface. Create a func “foo” that
has a value of type error as a parameter. Create a value of type “customErr” and pass it into
“foo”. If you need a hint, here is one: https://go.dev/play/p/L5QWV8-p11

	solution:
		https://play.golang.org/p/ixeowY2fd2
	assertion
		https://play.golang.org/p/pbl2kCYsM0
	conversion
		https://play.golang.org/p/1ldiBdkdzA
video: 178
*/

package main

import "fmt"

type customErr struct {
	info string
}

func main() {
	c := customErr{
		info: "I really don't know if you're supposed to touch that",
	}
	foo(c)
}
func (c customErr) Error() string {
	return fmt.Sprintf("here is the customError: %v\n", c.info)
}

func foo(err error) {
	fmt.Println(err)
}
