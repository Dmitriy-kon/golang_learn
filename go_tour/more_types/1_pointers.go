package main

import "fmt"

func main() {
	i, j := 32, 1020

	p := &i
	fmt.Printf("%p\n", p)
	fmt.Println(*p, j)
	*p /= 10
	fmt.Println(*p, j)
}
