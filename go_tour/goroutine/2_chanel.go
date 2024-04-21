package main

import (
	"fmt"
	"math/rand/v2"
	"slices"
	"sync"
)

func create_array(n int) []int {
	arr := make([]int, n)

	for i := range n {
		arr[i] = rand.IntN(200)
	}
	slices.Sort(arr)
	return arr
}

func sum(s []int, c chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	sum := 0

	for _, v := range s {
		sum += v
	}
	c <- sum
}

func main() {
	arr1 := create_array(8)
	var wg sync.WaitGroup
	// arr2 := create_array(6)

	c := make(chan int)
	wg.Add(3)

	go sum(arr1[:len(arr1)/2], c, &wg)
	go sum(arr1[len(arr1)/2:], c, &wg)
	go sum(arr1, c, &wg)

	x, y, z := <-c, <-c, <-c
	wg.Wait()

	close(c)

	fmt.Println(x, y, x+y)
	fmt.Printf("array sum from z = %d", z)

}
