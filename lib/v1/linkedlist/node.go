package linkedlist

import "fmt"

// Node represents a node in a single or double linked list.
type Node struct {
	data int
	next *Node
	prev *Node
}

// PrintSingleLL prints a node.
func (n *Node) PrintSingleLL() {
	fmt.Printf("| %d | %+v |\n", n.data, n.next)
}

func (n *Node) PrintDoubleLL() {
	fmt.Printf("| %+v | %d | %+v |\n", n.prev, n.data, n.next)
}
