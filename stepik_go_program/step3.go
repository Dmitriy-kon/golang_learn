package main

import "fmt"

func main() {
	a := []int{1, 2, 3}
	b := make([]int, 3, 3)
	n := copy(b, a)

	a[1] = 22
	fmt.Printf("a = %v, а тип %T \n", a, a)    // a = [1 22 3]
	fmt.Printf("b = %v, а тип %T \n", b, b)    // b = [1 2 3]
	fmt.Printf("Скопировано %d элемента\n", n) // Скопировано 3 элемента
}
