package linkedlist

import "fmt"

type Node struct {
	data int
	next *Node
}

type Single struct {
	size int
	head *Node
}

func NewSingleLinkedList() *Single {
	return &Single{head: nil}
}

// Length returns the length of the list.
func (s *Single) Length() int {
	return s.size
}

// IsEmpty checks if the list is empty or not.
func (s *Single) IsEmpty() bool {
	return s.Length() == 0
}

// updateLength updates the size of the list.
func (s *Single) updateLength() {
	s.size++
}

// Insert inserts the data to the end of the list
func (s *Single) Insert(data int) {
	node := &Node{
		data: data,
		next: nil,
	}

	// First element
	if s.head == nil {
		s.head = node
		s.updateLength()
		return
	}

	// traverse to last element
	lastNode := s.head
	for lastNode.next != nil {
		lastNode = lastNode.next
	}

	// insert
	lastNode.next = node
	s.updateLength()
}

// Print prints list to the console.
func (s *Single) Print() {
	if s.head == nil {
		fmt.Println("list is empty")
		return
	}

	node := s.head
	for {
		fmt.Println(node.data)
		if node.next == nil {
			break
		}
		node = node.next
	}
}
