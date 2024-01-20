package main

import "fmt"

// Обращение по индексам и изменение

func main() {
	var numbers [5]int = [5]int{1, 2, 3, 4, 5}

	fmt.Println(numbers[0]) // 1
	fmt.Println(numbers[4]) // 5

	numbers[0] = 87

	fmt.Println(numbers[0]) // 87

	a := [5]int{1, 2, 3, 4, 5}
	fmt.Println(a) // [1 2 3 4 5]

	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}
}
