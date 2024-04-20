package main

import "fmt"

func describe(i interface{}) {
	fmt.Printf("(value = %v and type = %T)\n", i, i)
}

func main() {
	var i interface{}
	describe(i)
	i = 42
	describe(i)
	i = "hello"
	describe(i)

}
