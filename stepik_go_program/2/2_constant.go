const (
	a int = 45
	b float32 = 3.3
)


package main

import (
	"fmt"
)

//  Также можно не указывать значение следующей константы по порядку (значение будет скопировано):
const(
	A int = 45
	B
	C float32 = 3.3
	D
)
func main() {
	fmt.Println(A, B, C, D)  // Вывод: 45 45 3.3 3.3
}
