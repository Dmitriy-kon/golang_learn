package main

import "fmt"

func main() {

	fmt.Print("Ivan", 27) // Ivan27 В первом случае один из объектов строка, поэтому пробел между объектами не ставится

	fmt.Println("Ivan", 27) // Ivan 27 Во втором случае мы используем метод Println()  поэтому пробел ставится в любом случае.

	fmt.Print(33, 27) // 33 27  третьем у нас нет строк - поэтому метод Print()   вставляет пробел между выводимыми объектами.

}
