package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["Answer"] = 48
	fmt.Printf("The value: %d\n", m["Answer"])

	elem := m["Answer"]
	fmt.Println(elem)

	delete(m, "Answer")
	fmt.Println(m)

	elem, ok := m["Answer"]
	if !ok {
		fmt.Println(ok)
	}
}
