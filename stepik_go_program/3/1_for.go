package main

import (
	"fmt"
)

// for [инициализация счетчика]; [условие]; [изменение счетчика]{
//     // действия
// }

func main() {
	// Инициализация переменно внутри for
	sum := 0
	for i := 1; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	// Нам необязательно указывать все условия при объявлении цикла. Например, можно вынести объявление переменной вовне:
	var i = 1
	for i < 10 {
		fmt.Println(i * i)
		i++
	}

	// Если цикл использует только условие, то его можно сократить следующим образом:
	var i = 1
	for i < 10 {
		fmt.Println(i * i)
		i++
	}

	// бесконечный цикл
	for {

	}

	var n int
	// считываем числа пока не будет введен 0
	for fmt.Scan(&n); n != 0; fmt.Scan(&n) {
		fmt.Println(n)
	}

}