package main

import "fmt"

type Vertex struct {
	X, Y int
}

func main() {
	fmt.Println(Vertex{1, 2})
	fmt.Println(Vertex{10, 20})
}
