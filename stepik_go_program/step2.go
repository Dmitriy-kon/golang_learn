package main

import (
	"fmt"
	"unicode"
)

func main() {
	var str string
	fmt.Scan(&str)

	runes := []rune(str)

	if len(runes) > 5 {
		for _, item := range runes {
			if !unicode.Is(unicode.Latin, item) && !unicode.IsDigit(item) {
				fmt.Println("Wrong password")
				return
			}
		}
		fmt.Println("Ok")
	} else {
		fmt.Println("Wrong password")
	}

}
