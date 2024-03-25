package main

import "fmt"

func main() {
	var s []int = []int{21, 12312, 123, 185}

	for i := range s {
		fmt.Println(i)
	}

}
