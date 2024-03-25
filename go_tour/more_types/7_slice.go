package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 4, 5, 6}

	// var s []int = primes[:3]
	// fmt.Println(s)

	// for i := range 6 {
	// 	println(primes[i])
	// }
	for _, item := range primes {
		fmt.Println(item)
	}
}
