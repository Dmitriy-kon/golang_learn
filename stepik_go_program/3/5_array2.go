package main

import (
	"fmt"
)

/*
Вместо явного указания длины массива мы можем указать символ многоточия
(...) в квадратных скобках, тогда длина массива будет определена Go в зависимости от количества указанных при инициализации значени
Не рекомендуются к использованию

d := [3]int{1: 12, 2: 24}
Такая форма позволяет явно указывать значения на индексах
*/

/*
Поскольку мы можем последовательно сравнить все элементы массива, мы можем сравнить и сами массивы:
Но при этом нужно учитывать, что сравнимы только массивы одного типа (массивы одинаковой длины, содержащие элементы одинакового типа).
*/
func main() {
	var a [3]int = [3]int{1, 2, 3}
	b := [3]int{1, 2, 3}

	c := [...]int{1, 2, 3, 4, 5}
	d := [3]int{1: 12, 2: 24}

	fmt.Println(a) // [1 2 3]
	fmt.Println(b) // [1 2 3]
	fmt.Println(c) // [1 2 3]
	fmt.Println(d) // [0 12 0]

	a := [3]int{1, 2, 3}
	b := [3]int{1, 2, 3}
	c := [3]int{3, 2, 1}

	fmt.Println(a == b) // true
	fmt.Println(a == c) // false
}