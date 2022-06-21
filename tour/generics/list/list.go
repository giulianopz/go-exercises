package main

import "fmt"

// List represents a singly-linked list that holds
// values of any type.
type List[T any] struct {
	next *List[T]
	val  T
}

func (l *List[int]) addNode(next *List[int]) {
	l.next = next
}

func main() {

	next := List[int]{nil,4}
	current := List[int]{nil, 3}
	current.addNode(&next)
	fmt.Println("current:",current)
	fmt.Println("next:",*current.next)
}
