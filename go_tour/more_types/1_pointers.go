package main

import "fmt"

func main() {
	i, j := 32, 1020

	p := &i
	fmt.Printf("%d\n", *p)

	newP := &j
	*newP /= 10

	fmt.Printf("%d", *newP)
}
