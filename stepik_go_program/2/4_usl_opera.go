package main

import "fmt"

func main() {
	var a int = 8
	var b int = 3
	var c bool = a == b
	fmt.Println(c) // false

	var a bool = true
	var b bool = !a // false
	var c bool = !b // true

	var b bool = 4 > 5 && 6 > 8   // false
	var c bool = 3 <= 5 && 10 > 8 // true

	var b bool = 4 > 5 || 6 > 8   // false
	var c bool = 3 == 5 || 10 > 8 // true
}
