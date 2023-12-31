package adt

import (
	"cmp"
	"fmt"
)

// Node represents a single node in linked list representation.
type Node[T any] struct {
	data T
	next *Node[T]
}

// NewNode creates an instance of Node with the data provided.
func NewNode[T any](data T) *Node[T] {
	return &Node[T]{
		data: data,
		next: nil,
	}
}

// Data returns the data from the Node.
func (n *Node[T]) Data() T {
	return n.data
}

// Display displays the node to the console.
func (n *Node[T]) Display() {
	fmt.Printf("%+v\n", n)
}

// TreeNode represents a node in a tree data structure.
type TreeNode[T cmp.Ordered] struct {
	data  T
	left  *TreeNode[T]
	right *TreeNode[T]
}

func NewTreeNode[T cmp.Ordered](data T) *TreeNode[T] {
	return &TreeNode[T]{
		data:  data,
		left:  nil,
		right: nil,
	}
}
