package main

import (
	"fmt"
)

func main() {
	var a, b int
	if a < b {
		fmt.Println("a меньше b")
	} else if a > b {
		fmt.Println("a больше b")
	} else {
		fmt.Println("a равно b")
	}
}
