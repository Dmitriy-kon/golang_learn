package main

import (
	"fmt"
)

func main() {
	var str, newstr string
	fmt.Scan(&str)

	for idx, item := range str {
		if idx%2 != 0 {
			newstr += string(item)
		}
	}
	fmt.Println(newstr)

}
