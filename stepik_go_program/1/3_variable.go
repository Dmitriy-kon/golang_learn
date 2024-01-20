package main

// var name type = "some" способ создания переменной
// одновременное определение нескольких переменных var a, b, c string

// Символы
// Символы пишутся в одинарных скобках '' и имеют int тип, к примеру int32

//Когда объявляется переменная, она автоматически содержит значение по умолчанию для своего типа:
// 0 для int, 0.0 для float, false для bool, пустая строка для string, nil для указателя и т.д.

func main() {
	var hello string = "Hello world"
	var _min int = 2
	var symbol int32 = 'C'

	println(string(symbol))
	println(hello)
	println(_min)

}
