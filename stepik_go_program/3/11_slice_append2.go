package main

import "fmt"

/*
И так, мы создали массив из 10 элементов типа int,
а затем создали срез на его элементы 5-7 (значения элементов в примере соответствуют их индексам).
Таким образом длина среза составляет 3, а емкость 5
*/

func main() {
	baseArray := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("Базовый массив: %v\n", baseArray)

	baseSlice := baseArray[5:8]
	fmt.Printf(
		"Срез, основанный на базовом массиве длиной %d и емкостью %d: %v\n",
		len(baseSlice),
		cap(baseSlice),
		baseSlice,
	)
	pointer := fmt.Sprintf("%p", baseSlice)

	baseSlice = append(baseSlice, 10)
	fmt.Printf("Массив: %v\n", baseArray)
	fmt.Printf("Срез длиной %d и емкостью %d: %v\n", len(baseSlice), cap(baseSlice), baseSlice)
	fmt.Println(pointer == fmt.Sprintf("%p", baseSlice))

	/*
		Мы видим, что базовый массив не изменился, а наш срез теперь ссылается на другой массив и имеет емкость больше длины.
		Почему так произошло? При добавлении элементов в срез Go проверяет, достаточно ли емкости среза для добавления новых элементов
		в срез (т.е. есть ли еще место в массиве, на котором основан срез).
		Если емкости не достаточно, то создается новый срез, основанный на массиве большего объема, в который
		копируются все элементы из старого среза, а также добавляются новые элементы.
	*/
	baseSlice = append(baseSlice, 10, 12, 14)
	fmt.Printf("Массив: %v\n", baseArray)
	fmt.Printf("Срез длиной %d и емкостью %d: %v\n", len(baseSlice), cap(baseSlice), baseSlice)
	fmt.Println(pointer == fmt.Sprintf("%p", baseSlice))

}
