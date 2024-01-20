package main

import (
	"fmt"
)

// iota идентификатор Go используется в объявлениях констант для упрощения определений увеличивающихся чисел.
// Сделаем дни недели с использованием iota - теперь это выглядит проще (особенно если много данных):

const (
	Sunday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	_ // пропускаем значение
	Add
)

const (
	u         = iota * 42 // u == 0 (индекс - 0, поэтому 0 * 42 = 0)
	v float64 = iota * 42 // v == 42.0 (индекс - 1, поэтому 1.0 * 42 = 42.0)
	w         = iota * 42 // w == 84  (индекс - 2, поэтому 2 * 42 = 84)
)

// переменные ни в одном блоке const, поэтому индекс не увеличился
const x = iota // x == 0
const y = iota // y == 0

func main() {
	fmt.Println(Sunday)   // вывод 0
	fmt.Println(Saturday) // вывод 6
	fmt.Println(Add)      // вывод 8
}
