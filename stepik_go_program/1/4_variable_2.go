package main

import "fmt"

// краткий способ объявлегие переменной
// a := 5
// var a int = 5
// или же var a = 5

// Форма объявление нескольких переменных

func main() {
	var (
		name string = "Dima"
		age  int    = 23
	)

	fmt.Println(name)
	fmt.Println(age)
}
