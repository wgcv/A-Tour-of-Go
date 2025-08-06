package main
package main

import (
	"fmt"
	)
// List represents a singly-linked list that holds
// values of any type.
type List[T any] struct {
	next *List[T]
	val  T
}

func (v *List[T]) Append(value T) {
	curr := v
	for curr.next != nil {
		curr = curr.next
	}
	curr.next = &List[T]{val: value}
}

func (v *List[T]) Print() {
	curr := v
	for curr != nil {
		fmt.Println(curr.val)
		curr = curr.next
	}
}

func main() {
	v := &List[int]{val: 1}
	v.Append(2)
	v.Append(3)

	v.Print()

}
