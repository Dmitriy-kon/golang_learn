package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 4, 5, 6}

	var s []int = primes[1:3]
	fmt.Println(s)
}