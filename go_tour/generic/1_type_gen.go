package main

import "fmt"

type Numbers interface {
	int | float64 | uint8
}

// Index returns the index of x in s, or -1 if not found.
func Index[T comparable](s []T, x T) int {
	for i, v := range s {
		// v and x are type T, which has the comparable
		// constraint, so we can use == here.
		if v == x {
			return i
		}
	}
	return -1
}

func Sum[T Numbers](arr []T) T {
	var res T
	for _, i := range arr {
		res = res + i
	}
	return res
}

func main() {
	// Index works on a slice of ints
	si := []int{10, 20, 15, -10}
	si2 := []uint8{10, 20, 15, 150}
	fmt.Println(Index(si, 15))
	fmt.Println(Sum(si))
	fmt.Println(Sum(si2))

	// Index also works on a slice of strings
	ss := []string{"foo", "bar", "baz"}
	fmt.Println(Index(ss, "hello"))
}
