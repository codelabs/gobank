package adt

import "fmt"

// Node represents a single node in linked list representation.
type Node[T any] struct {
	data T
	next *Node[T]
}

func (n *Node[T]) Data() T {
	return n.data
}

// Display displays the node to the console.
func (n *Node[T]) Display() {
	fmt.Printf("%+v\n", n)
}
