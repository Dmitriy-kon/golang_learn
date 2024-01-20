package main

import (
	"fmt"
	"math"
)

func main() {
	a := 6
	b := 2
	if a < b {
		fmt.Println("a меньше, чем b")
	}

	// Краткая запись if, тернарный

	var lim float64 = 37
	var (
		x float64 = 6
		y float64 = 2
	)

	if v := math.Pow(x, y); v < lim {
		fmt.Printf("%.2f", math.Round(v))
	}
}
