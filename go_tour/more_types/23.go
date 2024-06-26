package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	a, b := 0, 1

	return func() int {
		res := a
		a, b = b, a+b
		return res
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 50; i++ {
		fmt.Printf("%d ", f())
	}
}
