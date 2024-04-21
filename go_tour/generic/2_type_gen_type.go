package main

import "fmt"

// List represents a singly-linked list that holds
// values of any type.
type List[T any] struct {
	next *List[T]
	val  T
}

func main() {
	root := &List[string]{val: "string", next: nil}
	iter_l := root

	for i := range 10 {
		iter_l.next = &List[string]{val: fmt.Sprintf("string number %d", i), next: nil}
		iter_l = iter_l.next
	}

	iter_l2 := root
	for iter_l2 != nil {
		fmt.Println(iter_l2.val)
		iter_l2 = iter_l2.next
	}

}
