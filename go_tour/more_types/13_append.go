package main

import (
	"fmt"
	"slices"
)

func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(x), cap(x), x)
}

func main() {
	var s []int
	printSlice(s)

	s = append(s, 20)
	printSlice(s)

	s = append(s, 12, 41, 765)
	printSlice(s)
	s = append(s, 7, 0, -2)
	slices.Sort(s)
	printSlice(s)

}
