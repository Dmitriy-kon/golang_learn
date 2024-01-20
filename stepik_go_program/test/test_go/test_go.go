package main

import (
	"fmt"
	"strings"
	"unicode"
)

func latinLetters(s string) string {
	latstring := &strings.Builder{}

	for _, item := range s {
		if unicode.Is(unicode.Latin, item) {
			latstring.WriteString(string(item))
		}
	}
	return latstring.String()
}

func main() {
	fmt.Println(latinLetters("Привет мир World"))
}
