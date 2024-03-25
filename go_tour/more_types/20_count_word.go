package main

import (
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int {
	words := strings.Fields(s)

	m := make(map[string]int)

	for _, word := range words {
		// _, ok := m[word]
		// if !ok {
		// 	m[word] = 1
		// } else {
		// 	m[word] += 1
		// }
		m[word]++

	}
	return m
}

func main() {
	// wc.Test(WordCount)
	res := WordCount("some string some")
	fmt.Println(res)
}
