package main

import (
	"fmt"
)

type Vertex struct {
	Lat, Long float64
}

// var m map[string]Vertex

func main() {
	m := make(map[string]Vertex)
	m["Some Key"] = Vertex{
		23.12312312, -32.123123,
	}
	m["Another key"] = Vertex{
		12.12, -121.22,
	}
	m["Another key2"] = Vertex{
		12.12, -121.22,
	}

	for k, v := range m {
		fmt.Printf("%s -> %v\n", k, v)
	}

	fmt.Println(m)
}
