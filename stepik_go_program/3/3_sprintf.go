package main

/*
Также есть функция Sprintf() которая работает как и Printf(),
за исключением того что она ничего не печатает, а возвращает результат форматирования, рассмотрим пример:
*/

/*

Оператор	Формат	Пример кода	Результат
\\
Обратная косая черта

(это пример экранирования, а не управляющего символа)

fmt.Printf("\\")
\


\"
Двойные кавычки

(это пример экранирования, а не управляющего символа)

fmt.Printf("\"")
"
\f	подача страницы
fmt.Printf("123\f456\f789")

\v	вертикальный таб
fmt.Printf("\vf\v")



f

\r	возврат каретки
fmt.Printf("\r Input ")
fmt.Scan(&a)



\b	возврат (backspace U+0008)
fmt.Printf("123\b")
fmt.Scan(&a)



\t	Табуляция (горизонтальный отступ)
fmt.Printf("\ttest")


        test
\n	Перевод строки
fmt.Printf("test\ntest")


test
test


*/
import (
	"fmt"
)

func main() {
	var a float64 = 100.123456789
	result := fmt.Sprintf("%.2f", a)
	fmt.Printf("%q", result) // вывод: "100.12"
	// result будет типа string
}
