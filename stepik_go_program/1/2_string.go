package main

import "fmt"

// string() переводит цифру в строку, так же и срез
// "hello"[1] по индексу выдает номер буквы. Для получения буквы добавляем в string()
// "hello"[1:2] синтаксис среза, выдает строку

func main() {
	println("sometext")
	fmt.Println("SOMETEXT"[0])
	fmt.Println("SOMETEXT"[0:4])
}
