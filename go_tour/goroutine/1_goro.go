package main

import (
	"fmt"
	"sync"
	"time"
)

func say(s string, t int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		time.Sleep(time.Duration(t) * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	var wg sync.WaitGroup

	wg.Add(2)
	go say("world", 1500, &wg)
	go say("Привки", 1500, &wg)
	wg.Wait()
}
