package main

import (
	"fmt"
	"slices"
)

func main() {
	var s []int = []int{21, 12312, 123, 185}
	slices.Sort(s)

	for item := range len(s) {
		fmt.Println(s[item])
	}

}
