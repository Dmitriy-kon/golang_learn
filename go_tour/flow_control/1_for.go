package main

import "fmt"

func main() {
	sum := 0
	x := 1

	for i := 0; i < 10; i++ {
		sum += 1
	}

	for ; x < 10; x++ {
		x += 1
	}
	fmt.Println(x)
	fmt.Println(sum)

	for sum < 20 {
		sum++
	}
	fmt.Println(sum)

	for {
		if sum <= 1 {
			break
		}
		sum--
	}

	fmt.Println(sum)
}
