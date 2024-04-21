package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 5) // Создаем буферизированный канал с размером буфера 5

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Printf("Отправляемое значение: %d\n", i)
			ch <- i                 // Отправляем значения в канал
			time.Sleep(time.Second) // Имитируем задержку
		}
		close(ch) // Закрываем канал после отправки всех значений
	}()

	// Даем время горутине начать работу
	// time.Sleep(time.Second * 2)

	for i := range ch {
		fmt.Printf("Принятое значение: %d\n", i)
		// Даем время для чтения из канала, чтобы показать, как работает буфер
		time.Sleep(time.Second * 2)
	}
}
