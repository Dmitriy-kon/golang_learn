package main

import (
	"fmt"
	"sync"
)

func fibonacci(n int, c chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	c := make(chan int, 3)

	go fibonacci(10, c, &wg)

	go func() {
		defer wg.Done()
		for i := range c {
			fmt.Println(i)
		}
	}()

	wg.Wait()
}
