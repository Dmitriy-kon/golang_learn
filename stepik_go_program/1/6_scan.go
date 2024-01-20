package main

import "fmt"

// Scan позволяет вводить запрос с клавиатуры в программу
// &var ссылка на переменную. Вернее адрес
// чтение нескольких переменных с консоли fmt.Scan(&a, &b, &c)

func main() {
	var name string
	var age int8

	println("Введи имя:")
	fmt.Scan(&name)

	println("Введи возраст:")
	fmt.Scan(&age)

	println(name, age)

}
